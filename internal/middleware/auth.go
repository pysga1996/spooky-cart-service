package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
	"github.com/square/go-jose"
	"github.com/thanh-vt/splash-inventory-service/internal"
	"github.com/thanh-vt/splash-inventory-service/internal/constant"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getJSONWebKeySetFromCache() (*jose.JSONWebKeySet, bool) {
	val, err := internal.Redis.Get(internal.RedisCtx, constant.JwksCacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, false
		}
		panic(err)
	}
	jsonWebKeySet := new(jose.JSONWebKeySet)
	if err = json.Unmarshal([]byte(val), &jsonWebKeySet); err != nil {
		return nil, false
	}

	return jsonWebKeySet, true
}

func fetchJSONWebKeySet() *jose.JSONWebKeySet {
	var resp *http.Response
	var err error
	var jwksUrl = os.Getenv("JWKS_URL")
	log.Printf("Fetchng JWKS from %s", jwksUrl)

	if resp, err = internal.HttpClient.Get(jwksUrl); err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			panic(err)
		}
	}(resp.Body)
	if resp.StatusCode != 200 {
		panic(fmt.Errorf("failed request, status: %d", resp.StatusCode))
	}

	jsonWebKeySet := new(jose.JSONWebKeySet)
	if err = json.NewDecoder(resp.Body).Decode(jsonWebKeySet); err != nil {
		panic(err)
	}

	return jsonWebKeySet
}

func getJSONWebKey(keyId string) *jose.JSONWebKey {
	var val []byte
	var err error
	jwks, exists := getJSONWebKeySetFromCache()
	if !exists {
		jwks = fetchJSONWebKeySet()
		if val, err = json.Marshal(jwks); err != nil {
			panic(err)
		}
		if _, err = internal.Redis.Set(internal.RedisCtx, constant.JwksCacheKey,
			string(val), 24*time.Hour).Result(); err != nil {
			panic(err)
		}
	}
	var keys []jose.JSONWebKey

	if keyId == "" {
		keys = jwks.Keys
	} else {
		keys = jwks.Key(keyId)
	}
	if len(keys) == 0 {
		panic(fmt.Errorf("JWK is not found: %s", keyId))
	}
	return &keys[0]
}

func HandleToken(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// sample token string taken from the New example
		authHeader := r.Header.Get("Authorization")

		tokenArr := strings.Split(authHeader, " ")
		if len(tokenArr) < 2 {
			next.ServeHTTP(w, r)
			return
		}
		tokenStr := tokenArr[1]
		jwk := getJSONWebKey("")

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return jwk.Certificates[0].PublicKey, nil
		})

		if err != nil {
			Unauthorized(w, r, err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), constant.UID, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			BadRequest(w, r, errors.New("invalid or expired jwt"))
			return
		}

	}

	return http.HandlerFunc(fn)

}

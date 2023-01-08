package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thanh-vt/spooky-cart-service/internal/constant"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/s12v/go-jwks"
)

var jwksSource = jwks.NewWebSource(os.Getenv("JWKS_URL"))

var jwksClient = jwks.NewDefaultClient(
	jwksSource,
	time.Hour,    // Refresh keys every 1 hour
	12*time.Hour, // Expire keys after 12 hours
)

func HandleToken(c *gin.Context) {

	// sample token string taken from the New example
	authHeader := c.GetHeader("Authorization")

	tokenArr := strings.Split(authHeader, " ")
	if len(tokenArr) < 2 {
		return
	}
	tokenStr := tokenArr[1]
	jwk, err := jwksClient.GetSignatureKey(os.Getenv("JWKS_ID"))
	if err != nil {
		BadRequest(c, err)
		return
	}
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
		Unauthorized(c, err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set(constant.UID, claims["email"])
	} else {
		BadRequest(c, errors.New("invalid or expired jwt"))
		return
	}
}

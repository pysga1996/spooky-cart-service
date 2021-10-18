package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	pkgErr "github.com/pysga1996/spooky-cart-service/error"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/s12v/go-jwks"
	"github.com/square/go-jose"
)

var publicKey = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAxxxxxxf2iF+20xHTZ4jTUBzYmikBuUsm0839T5SDmwEquTB\nfQIDAQAB\n-----END PUBLIC KEY-----\n"

var jwksSource = jwks.NewWebSource("https://lemur-6.cloud-iam.com/auth/realms/plate/protocol/openid-connect/certs")

var jwk *jose.JSONWebKey

var jwksClient = jwks.NewDefaultClient(
	jwksSource,
	time.Hour,    // Refresh keys every 1 hour
	12*time.Hour, // Expire keys after 12 hours
)

func Handle(c *gin.Context) {

	// sample token string taken from the New example
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Split(authHeader, " ")[0]

	jwk, err := jwksClient.GetSignatureKey("0txynZO10IrruiuyZcUOZUBIRbnECjkHXWUcX_xbG5Y")
	if err != nil {
		pkgErr.BadRequest(c, err)
		return
	}
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwk.Certificates[0].Raw, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["email"])
		c.Set("username", claims["email"])
	} else {
		fmt.Println(err)
	}
	c.Next()
}

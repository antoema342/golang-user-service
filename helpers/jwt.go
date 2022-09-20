package helpers

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type AuthCustomClaims struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	User bool      `json:"user"`
	jwt.StandardClaims
}

func GenerateToken(id uuid.UUID, name string, isUser bool) string {
	claims := &AuthCustomClaims{
		id,
		name,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    name,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte("iniRAHASIA"))
	if err != nil {
		panic(err)
	}
	return t
}

var jwtKey = []byte("iniRAHASIA")

func MiddlewareJWTAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Request header auth Empty",
			})
			c.Abort()
			return
		}
		if !strings.Contains(authorizationHeader, "Bearer") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		claims := &AuthCustomClaims{}

		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid token",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}
		if !tkn.Valid {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}
		c.Set("decoded", claims)
		c.Next()
	}
	// ...

}

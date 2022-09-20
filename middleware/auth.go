package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"strings"
	"upload_file_handler/util"
)

var JwtKey = []byte("123456")

type authHeader struct {
	token string `header:"Authorization"`
}

func AuthorizeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}

		// bind Authorization Header to h and check for validation errors
		if err := c.ShouldBindHeader(&h); err != nil {
			util.SendError(c, 403, err)
			c.Abort()
			return
		}
		token := c.Request.Header["Authorization"]
		jwtToken := strings.Split(token[0], " ")[1]
		parsedToken, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("authorization failed")
			}
			return JwtKey, nil
		})

		if err != nil {
			util.SendError(c, 403, err)
			c.Abort()
			return
		}
		if parsedToken.Valid == false {
			util.SendError(c, 403, errors.New("token is invalid"))
			c.Abort()
			return
		}
		payload, ok := parsedToken.Claims.(jwt.MapClaims)
		if ok == false {
			util.SendError(c, 403, errors.New("parse token payload failed"))
			c.Abort()
			return
		}

		c.Set("user", payload)
		c.Set("userId", payload["userId"])
		c.Next()
	}

}

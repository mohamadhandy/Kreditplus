package middleware

import (
	"fmt"
	"kredit-plus/helper"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	IDKonsumen int     `json:"id_konsumen"`
	Email      string  `json:"email"`
	Name       string  `json:"name"`
	Role       string  `json:"role"`
	Gaji       float64 `json:"gaji"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			data := helper.Response{
				Data:       nil,
				Message:    err.Error(),
				StatusCode: http.StatusUnprocessableEntity,
			}
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, data)
			return
		}
		if _, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, "unauthorized err")
		}
	}
}

func ExtractRoleFromToken(tokenString string) string {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		fmt.Println("error extract token name: " + err.Error())
		return err.Error()
	}

	// Extract the user ID from the token claims
	if res, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return res.Role
	} else {
		return "Token Not Valid"
	}
}

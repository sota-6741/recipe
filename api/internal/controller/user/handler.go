package user

import (
	"fmt"
	"net/http"

	"recipe/api/internal/middleware/auth0"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	subToUsers = map[string]User{
		"auth0|61a8178b21127500715968e2": {
			Name: "kourin",
			Age:  15,
		},
	}
)

func getUser(sub string) *User {
	user, ok := subToUsers[sub]
	if !ok {
		return nil
	}
	return &user
}

type UserHandler struct{}

func (h *UserHandler) Me(c *gin.Context) {
	token := auth0.GetJWT(c.Request.Context())
	fmt.Printf("jwt %+v\n", token)

	if token == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid claims"})
		return
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "sub claim missing"})
		return
	}

	user := getUser(sub)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

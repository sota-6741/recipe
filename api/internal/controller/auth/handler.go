package auth

import "github.com/gin-gonic/gin"

type AuthHandler struct{}

func (h *AuthHandler) GoogleLogin(c *gin.Context) {}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

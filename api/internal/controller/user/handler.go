package user

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func (h *UserHandler) Me(c *gin.Context) {}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

package user

import (
	"net/http"

	"recipe/api/internal/middleware/auth0"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	// TODO: UserRepositoryを注入
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// Me は認証済みユーザーの情報を取得する
// GET /user/me
func (h *UserHandler) Me(c *gin.Context) {
	userID := auth0.GetUserID(c.Request.Context())
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です"})
		return
	}

	// TODO: DBからユーザー情報を取得
	// user, err := h.userRepo.FindByID(userID)
	// if err != nil {
	//     c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
	//     return
	// }

	c.JSON(http.StatusOK, MeResponse{
		ID: userID,
	})
}

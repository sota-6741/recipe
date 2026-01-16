package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	// TODO: UserRepository, JWTServiceを注入
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// GoogleLogin はGoogleアカウントでログインする
// POST /auth/google
func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	var req GoogleLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "idTokenは必須です"})
		return
	}

	// TODO: Google ID トークンを検証
	// googleUser, err := verifyGoogleIDToken(req.IDToken)
	// if err != nil {
	//     c.JSON(http.StatusUnauthorized, gin.H{"error": "無効なIDトークンです"})
	//     return
	// }

	// TODO: ユーザーを検索または作成
	// user, err := h.userRepo.FindOrCreateByGoogleSub(googleUser.Sub)
	// if err != nil {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザー処理に失敗しました"})
	//     return
	// }

	// TODO: JWTを生成
	// jwt, err := h.jwtService.Generate(user.ID)
	// if err != nil {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT生成に失敗しました"})
	//     return
	// }

	// 仮実装: DBとJWTサービス実装後に置き換え
	c.JSON(http.StatusOK, GoogleLoginResponse{
		JWT: "TODO: generate jwt",
		User: UserResponse{
			ID: "TODO: user id",
		},
	})
}

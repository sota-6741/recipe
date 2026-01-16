package auth

// GoogleLoginRequest はGoogleログインのリクエスト
type GoogleLoginRequest struct {
	IDToken string `json:"idToken" binding:"required"`
}

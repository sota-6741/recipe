package auth

// GoogleLoginResponse はGoogleログインのレスポンス
type GoogleLoginResponse struct {
	JWT  string       `json:"jwt"`
	User UserResponse `json:"user"`
}

// UserResponse はユーザー情報のレスポンス
type UserResponse struct {
	ID string `json:"id"`
}

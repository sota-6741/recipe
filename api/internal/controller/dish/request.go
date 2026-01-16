package dish

// CreateDishRequest は料理作成のリクエスト
type CreateDishRequest struct {
	Name     string `json:"name" binding:"required"`
	NameKana string `json:"nameKana" binding:"required"`
}

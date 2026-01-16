package recipe

// CreateRecipeRequest はレシピ作成のリクエスト
type CreateRecipeRequest struct {
	DishID      string   `json:"dishId" binding:"required"`
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
}

// UpdateRecipeRequest はレシピ更新のリクエスト
type UpdateRecipeRequest struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
}

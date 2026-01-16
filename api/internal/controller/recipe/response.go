package recipe

import "time"

// RecipeResponse はレシピ情報のレスポンス
type RecipeResponse struct {
	ID          string    `json:"id"`
	DishID      string    `json:"dishId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Ingredients []string  `json:"ingredients"`
	Steps       []string  `json:"steps"`
	CreatedAt   time.Time `json:"createdAt"`
}

// RecipeListResponse はレシピ一覧のレスポンス
type RecipeListResponse struct {
	Recipes []RecipeResponse `json:"recipes"`
}

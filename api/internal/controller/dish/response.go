package dish

import "time"

// DishResponse は料理情報のレスポンス
type DishResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	NameKana  string    `json:"nameKana"`
	CreatedAt time.Time `json:"createdAt"`
}

// DishListResponse は料理一覧のレスポンス
type DishListResponse struct {
	Dishes []DishResponse `json:"dishes"`
}

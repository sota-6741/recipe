package domain

import "context"

// UserRepository はユーザーの永続化を担当する
type UserRepository interface {
	// FindByID はIDでユーザーを取得する
	FindByID(ctx context.Context, id string) (*User, error)
	// FindByGoogleSub はGoogleSubでユーザーを取得する
	FindByGoogleSub(ctx context.Context, googleSub string) (*User, error)
	// Create は新しいユーザーを作成する
	Create(ctx context.Context, user *User) error
}

// DishRepository は料理の永続化を担当する
type DishRepository interface {
	// FindByID はIDで料理を取得する
	FindByID(ctx context.Context, id string) (*Dish, error)
	// FindByUserID はユーザーIDで料理一覧を取得する（50音順）
	FindByUserID(ctx context.Context, userID string) ([]*Dish, error)
	// Create は新しい料理を作成する
	Create(ctx context.Context, dish *Dish) error
	// Delete は料理を削除する（レシピもカスケード削除）
	Delete(ctx context.Context, id string) error
}

// RecipeRepository はレシピの永続化を担当する
type RecipeRepository interface {
	// FindByID はIDでレシピを取得する
	FindByID(ctx context.Context, id string) (*Recipe, error)
	// FindByDishID は料理IDでレシピ一覧を取得する
	FindByDishID(ctx context.Context, dishID string) ([]*Recipe, error)
	// Create は新しいレシピを作成する
	Create(ctx context.Context, recipe *Recipe) error
	// Update はレシピを更新する
	Update(ctx context.Context, recipe *Recipe) error
	// Delete はレシピを削除する
	Delete(ctx context.Context, id string) error
}

package usecase

import (
	"context"
	"errors"

	"recipe/api/internal/domain"
)

var (
	ErrRecipeNotFound  = errors.New("recipe not found")
	ErrRecipeForbidden = errors.New("recipe access forbidden")
)

// RecipeUsecase はレシピに関するユースケース
type RecipeUsecase struct {
	recipeRepo domain.RecipeRepository
	dishRepo   domain.DishRepository
}

// NewRecipeUsecase はRecipeUsecaseを生成する
func NewRecipeUsecase(recipeRepo domain.RecipeRepository, dishRepo domain.DishRepository) *RecipeUsecase {
	return &RecipeUsecase{
		recipeRepo: recipeRepo,
		dishRepo:   dishRepo,
	}
}

// ListByDish は料理に紐づくレシピ一覧を取得する
func (u *RecipeUsecase) ListByDish(ctx context.Context, userID, dishID string) ([]*domain.Recipe, error) {
	// 料理の所有者確認
	dish, err := u.dishRepo.FindByID(ctx, dishID)
	if err != nil {
		return nil, ErrDishNotFound
	}

	if !dish.IsOwnedBy(userID) {
		return nil, ErrDishForbidden
	}

	return u.recipeRepo.FindByDishID(ctx, dishID)
}

// Get はレシピを取得する
func (u *RecipeUsecase) Get(ctx context.Context, userID, recipeID string) (*domain.Recipe, error) {
	recipe, err := u.recipeRepo.FindByID(ctx, recipeID)
	if err != nil {
		return nil, ErrRecipeNotFound
	}

	// 料理の所有者確認
	dish, err := u.dishRepo.FindByID(ctx, recipe.DishID())
	if err != nil {
		return nil, ErrDishNotFound
	}

	if !dish.IsOwnedBy(userID) {
		return nil, ErrRecipeForbidden
	}

	return recipe, nil
}

// CreateSiteRecipe はサイトレシピを作成する
func (u *RecipeUsecase) CreateSiteRecipe(ctx context.Context, userID, dishID, title, url string) (*domain.Recipe, error) {
	// 料理の所有者確認
	dish, err := u.dishRepo.FindByID(ctx, dishID)
	if err != nil {
		return nil, ErrDishNotFound
	}

	if !dish.IsOwnedBy(userID) {
		return nil, ErrDishForbidden
	}

	recipe, err := domain.NewSiteRecipe(dishID, title, url)
	if err != nil {
		return nil, err
	}

	if err := u.recipeRepo.Create(ctx, recipe); err != nil {
		return nil, err
	}

	return recipe, nil
}

// CreateCustomRecipe は自作レシピを作成する
func (u *RecipeUsecase) CreateCustomRecipe(ctx context.Context, userID, dishID, title, markdown string) (*domain.Recipe, error) {
	// 料理の所有者確認
	dish, err := u.dishRepo.FindByID(ctx, dishID)
	if err != nil {
		return nil, ErrDishNotFound
	}

	if !dish.IsOwnedBy(userID) {
		return nil, ErrDishForbidden
	}

	recipe, err := domain.NewCustomRecipe(dishID, title, markdown)
	if err != nil {
		return nil, err
	}

	if err := u.recipeRepo.Create(ctx, recipe); err != nil {
		return nil, err
	}

	return recipe, nil
}

// Update はレシピを更新する
func (u *RecipeUsecase) Update(ctx context.Context, userID, recipeID, title string, url, markdown *string) (*domain.Recipe, error) {
	recipe, err := u.recipeRepo.FindByID(ctx, recipeID)
	if err != nil {
		return nil, ErrRecipeNotFound
	}

	// 料理の所有者確認
	dish, err := u.dishRepo.FindByID(ctx, recipe.DishID())
	if err != nil {
		return nil, ErrDishNotFound
	}

	if !dish.IsOwnedBy(userID) {
		return nil, ErrRecipeForbidden
	}

	// レシピを更新
	if err := recipe.Update(title, url, markdown); err != nil {
		return nil, err
	}

	if err := u.recipeRepo.Update(ctx, recipe); err != nil {
		return nil, err
	}

	return recipe, nil
}

// Delete はレシピを削除する
func (u *RecipeUsecase) Delete(ctx context.Context, userID, recipeID string) error {
	recipe, err := u.recipeRepo.FindByID(ctx, recipeID)
	if err != nil {
		return ErrRecipeNotFound
	}

	// 料理の所有者確認
	dish, err := u.dishRepo.FindByID(ctx, recipe.DishID())
	if err != nil {
		return ErrDishNotFound
	}

	if !dish.IsOwnedBy(userID) {
		return ErrRecipeForbidden
	}

	return u.recipeRepo.Delete(ctx, recipeID)
}

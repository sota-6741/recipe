package usecase

import (
	"context"
	"errors"

	"recipe/api/internal/domain"
)

var (
	ErrDishNotFound  = errors.New("dish not found")
	ErrDishForbidden = errors.New("dish access forbidden")
)

// DishUsecase は料理に関するユースケース
type DishUsecase struct {
	dishRepo domain.DishRepository
}

// NewDishUsecase はDishUsecaseを生成する
func NewDishUsecase(dishRepo domain.DishRepository) *DishUsecase {
	return &DishUsecase{
		dishRepo: dishRepo,
	}
}

// List はユーザーの料理一覧を取得する
func (u *DishUsecase) List(ctx context.Context, userID string) ([]*domain.Dish, error) {
	return u.dishRepo.FindByUserID(ctx, userID)
}

// Create は新しい料理を作成する
func (u *DishUsecase) Create(ctx context.Context, userID, name, nameKana string) (*domain.Dish, error) {
	dish, err := domain.NewDish(userID, name, nameKana)
	if err != nil {
		return nil, err
	}

	if err := u.dishRepo.Create(ctx, dish); err != nil {
		return nil, err
	}

	return dish, nil
}

// Delete は料理を削除する
func (u *DishUsecase) Delete(ctx context.Context, userID, dishID string) error {
	dish, err := u.dishRepo.FindByID(ctx, dishID)
	if err != nil {
		return ErrDishNotFound
	}

	if !dish.IsOwnedBy(userID) {
		return ErrDishForbidden
	}

	return u.dishRepo.Delete(ctx, dishID)
}

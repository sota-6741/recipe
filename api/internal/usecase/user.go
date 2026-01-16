package usecase

import (
	"context"

	"recipe/api/internal/domain"
)

// UserUsecase はユーザーに関するユースケース
type UserUsecase struct {
	userRepo domain.UserRepository
}

// NewUserUsecase はUserUsecaseを生成する
func NewUserUsecase(userRepo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

// GetByID はIDでユーザーを取得する
func (u *UserUsecase) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return u.userRepo.FindByID(ctx, id)
}

// FindOrCreateByGoogleSub はGoogleSubでユーザーを検索し、存在しなければ作成する
func (u *UserUsecase) FindOrCreateByGoogleSub(ctx context.Context, googleSub string) (*domain.User, error) {
	user, err := u.userRepo.FindByGoogleSub(ctx, googleSub)
	if err == nil {
		return user, nil
	}

	// ユーザーが存在しない場合は新規作成
	newUser, err := domain.NewUser(googleSub)
	if err != nil {
		return nil, err
	}

	if err := u.userRepo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}
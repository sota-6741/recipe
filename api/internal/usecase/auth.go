package usecase

import (
	"context"

	"recipe/api/internal/domain"
)

// AuthUsecase は認証に関するユースケース
type AuthUsecase struct {
	userRepo domain.UserRepository
}

// NewAuthUsecase はAuthUsecaseを生成する
func NewAuthUsecase(userRepo domain.UserRepository) *AuthUsecase {
	return &AuthUsecase{
		userRepo: userRepo,
	}
}

// GoogleLogin はGoogleログインを処理する
func (u *AuthUsecase) GoogleLogin(ctx context.Context, googleSub string) (*domain.User, error) {
	// 既存ユーザーを検索
	user, err := u.userRepo.FindByGoogleSub(ctx, googleSub)
	if err == nil {
		return user, nil
	}

	// 新規ユーザーを作成
	newUser, err := domain.NewUser(googleSub)
	if err != nil {
		return nil, err
	}

	if err := u.userRepo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

package domain

import (
"errors"
"time"

"github.com/google/uuid"
)

// エラー定義
var (
ErrInvalidUserID        = errors.New("invalid user id")
ErrInvalidUserGoogleSub = errors.New("invalid user google sub")
)

// ドメインモデル
type User struct {
	id        userID
	googleSub userGoogleSub
	createdAt time.Time
}

// バリューオブジェクト
type userID struct{ value string }
type userGoogleSub struct{ value string }

// コンストラクタ
func NewUser(googleSub string) (*User, error) {
	id, err := newUserID(uuid.New().String())
	if err != nil {
		return nil, err
	}

	sub, err := newUserGoogleSub(googleSub)
	if err != nil {
		return nil, err
	}

	return &User{
		id:        *id,
		googleSub: *sub,
		createdAt: time.Now(),
	}, nil
}

// ReconstructUser はDBから取得したデータでUserを再構築する
func ReconstructUser(id, googleSub string, createdAt time.Time) *User {
	return &User{
		id:        userID{value: id},
		googleSub: userGoogleSub{value: googleSub},
		createdAt: createdAt,
	}
}

// バリューオブジェクトのコンストラクタ
func newUserID(value string) (*userID, error) {
	v := &userID{value: value}
	if err := v.validate(); err != nil {
		return nil, err
	}
	return v, nil
}

func newUserGoogleSub(value string) (*userGoogleSub, error) {
	v := &userGoogleSub{value: value}
	if err := v.validate(); err != nil {
		return nil, err
	}
	return v, nil
}

// バリデーション
func (v *userID) validate() error {
	if v.value == "" {
		return ErrInvalidUserID
	}
	return nil
}

func (v *userGoogleSub) validate() error {
	if v.value == "" {
		return ErrInvalidUserGoogleSub
	}
	return nil
}

// ゲッター
func (u *User) ID() string           { return u.id.value }
func (u *User) GoogleSub() string    { return u.googleSub.value }
func (u *User) CreatedAt() time.Time { return u.createdAt }

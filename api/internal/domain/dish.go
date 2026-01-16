package domain

import (
"errors"
"time"

"github.com/google/uuid"
)

// エラー定義
var (
ErrInvalidDishID       = errors.New("invalid dish id")
ErrInvalidDishUserID   = errors.New("invalid dish user id")
ErrInvalidDishName     = errors.New("invalid dish name")
ErrInvalidDishNameKana = errors.New("invalid dish name kana")
)

// ドメインモデル
type Dish struct {
	id        dishID
	userID    dishUserID
	name      dishName
	nameKana  dishNameKana
	createdAt time.Time
}

// バリューオブジェクト
type dishID struct{ value string }
type dishUserID struct{ value string }
type dishName struct{ value string }
type dishNameKana struct{ value string }

// コンストラクタ
func NewDish(userID, name, nameKana string) (*Dish, error) {
	id, err := newDishID(uuid.New().String())
	if err != nil {
		return nil, err
	}

	uid, err := newDishUserID(userID)
	if err != nil {
		return nil, err
	}

	n, err := newDishName(name)
	if err != nil {
		return nil, err
	}

	nk, err := newDishNameKana(nameKana)
	if err != nil {
		return nil, err
	}

	return &Dish{
		id:        *id,
		userID:    *uid,
		name:      *n,
		nameKana:  *nk,
		createdAt: time.Now(),
	}, nil
}

// ReconstructDish はDBから取得したデータでDishを再構築する
func ReconstructDish(id, userID, name, nameKana string, createdAt time.Time) *Dish {
	return &Dish{
		id:        dishID{value: id},
		userID:    dishUserID{value: userID},
		name:      dishName{value: name},
		nameKana:  dishNameKana{value: nameKana},
		createdAt: createdAt,
	}
}

// バリューオブジェクトのコンストラクタ
func newDishID(value string) (*dishID, error) {
	v := &dishID{value: value}
	if err := v.validate(); err != nil {
		return nil, err
	}
	return v, nil
}

func newDishUserID(value string) (*dishUserID, error) {
	v := &dishUserID{value: value}
	if err := v.validate(); err != nil {
		return nil, err
	}
	return v, nil
}

func newDishName(value string) (*dishName, error) {
	v := &dishName{value: value}
	if err := v.validate(); err != nil {
		return nil, err
	}
	return v, nil
}

func newDishNameKana(value string) (*dishNameKana, error) {
	v := &dishNameKana{value: value}
	if err := v.validate(); err != nil {
		return nil, err
	}
	return v, nil
}

// バリデーション
func (v *dishID) validate() error {
	if v.value == "" {
		return ErrInvalidDishID
	}
	return nil
}

func (v *dishUserID) validate() error {
	if v.value == "" {
		return ErrInvalidDishUserID
	}
	return nil
}

func (v *dishName) validate() error {
	if v.value == "" {
		return ErrInvalidDishName
	}
	return nil
}

func (v *dishNameKana) validate() error {
	if v.value == "" {
		return ErrInvalidDishNameKana
	}
	return nil
}

// ゲッター
func (d *Dish) ID() string           { return d.id.value }
func (d *Dish) UserID() string       { return d.userID.value }
func (d *Dish) Name() string         { return d.name.value }
func (d *Dish) NameKana() string     { return d.nameKana.value }
func (d *Dish) CreatedAt() time.Time { return d.createdAt }

// ドメインロジック
func (d *Dish) IsOwnedBy(userID string) bool {
	return d.userID.value == userID
}

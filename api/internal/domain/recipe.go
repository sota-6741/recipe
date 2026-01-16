package domain

import (
"errors"
"time"

"github.com/google/uuid"
)

// エラー定義
var (
ErrInvalidRecipeID       = errors.New("invalid recipe id")
ErrInvalidRecipeDishID   = errors.New("invalid recipe dish id")
ErrInvalidRecipeTitle    = errors.New("invalid recipe title")
ErrInvalidRecipeType     = errors.New("invalid recipe type")
ErrInvalidRecipeURL      = errors.New("url is required for site recipe")
ErrInvalidRecipeMarkdown = errors.New("markdown is required for custom recipe")
)

// 定数
const (
RecipeTypeSite   = "site"
RecipeTypeCustom = "custom"
)

// ドメインモデル
type Recipe struct {
	id         recipeID
	dishID     recipeDishID
	title      recipeTitle
	recipeType recipeType
	url        recipeURL
	markdown   recipeMarkdown
	createdAt  time.Time
	updatedAt  time.Time
}

// バリューオブジェクト
type recipeID struct{ value string }
type recipeDishID struct{ value string }
type recipeTitle struct{ value string }
type recipeType struct{ value string }
type recipeURL struct{ value string }
type recipeMarkdown struct{ value string }

// コンストラクタ（サイトレシピ）
func NewSiteRecipe(dishID, title, url string) (*Recipe, error) {
	id, err := newRecipeID(uuid.New().String())
	if err != nil {
		return nil, err
	}

	did, err := newRecipeDishID(dishID)
	if err != nil {
		return nil, err
	}

	t, err := newRecipeTitle(title)
	if err != nil {
		return nil, err
	}

	rt, err := newRecipeType(RecipeTypeSite)
	if err != nil {
		return nil, err
	}

	u, err := newRecipeURL(url, RecipeTypeSite)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Recipe{
		id:         *id,
		dishID:     *did,
		title:      *t,
		recipeType: *rt,
		url:        *u,
		markdown:   recipeMarkdown{value: ""},
		createdAt:  now,
		updatedAt:  now,
	}, nil
}

// コンストラクタ（自作レシピ）
func NewCustomRecipe(dishID, title, markdown string) (*Recipe, error) {
	id, err := newRecipeID(uuid.New().String())
	if err != nil {
		return nil, err
	}

	did, err := newRecipeDishID(dishID)
	if err != nil {
		return nil, err
	}

	t, err := newRecipeTitle(title)
	if err != nil {
		return nil, err
	}

	rt, err := newRecipeType(RecipeTypeCustom)
	if err != nil {
		return nil, err
	}

	m, err := newRecipeMarkdown(markdown, RecipeTypeCustom)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Recipe{
		id:         *id,
		dishID:     *did,
		title:      *t,
		recipeType: *rt,
		url:        recipeURL{value: ""},
		markdown:   *m,
		createdAt:  now,
		updatedAt:  now,
	}, nil
}

// ReconstructRecipe はDBから取得したデータでRecipeを再構築する
func ReconstructRecipe(id, dishID, title, rType, url, markdown string, createdAt, updatedAt time.Time) *Recipe {
	return &Recipe{
		id:         recipeID{value: id},
		dishID:     recipeDishID{value: dishID},
		title:      recipeTitle{value: title},
		recipeType: recipeType{value: rType},
		url:        recipeURL{value: url},
		markdown:   recipeMarkdown{value: markdown},
		createdAt:  createdAt,
		updatedAt:  updatedAt,
	}
}

// バリューオブジェクトのコンストラクタ
func newRecipeID(value string) (*recipeID, error) {
	v := &recipeID{value: value}
	if err := v.validate(); err != nil {
		return nil, err
	}
	return v, nil
}

func newRecipeDishID(value string) (*recipeDishID, error) {
	v := &recipeDishID{value: value}
	if err := v.validate(); err != nil {
		return nil, err
	}
	return v, nil
}

func newRecipeTitle(value string) (*recipeTitle, error) {
	v := &recipeTitle{value: value}
	if err := v.validate(); err != nil {
		return nil, err
	}
	return v, nil
}

func newRecipeType(value string) (*recipeType, error) {
	v := &recipeType{value: value}
	if err := v.validate(); err != nil {
		return nil, err
	}
	return v, nil
}

func newRecipeURL(value string, rType string) (*recipeURL, error) {
	v := &recipeURL{value: value}
	if err := v.validate(rType); err != nil {
		return nil, err
	}
	return v, nil
}

func newRecipeMarkdown(value string, rType string) (*recipeMarkdown, error) {
	v := &recipeMarkdown{value: value}
	if err := v.validate(rType); err != nil {
		return nil, err
	}
	return v, nil
}

// バリデーション
func (v *recipeID) validate() error {
	if v.value == "" {
		return ErrInvalidRecipeID
	}
	return nil
}

func (v *recipeDishID) validate() error {
	if v.value == "" {
		return ErrInvalidRecipeDishID
	}
	return nil
}

func (v *recipeTitle) validate() error {
	if v.value == "" {
		return ErrInvalidRecipeTitle
	}
	return nil
}

func (v *recipeType) validate() error {
	if v.value != RecipeTypeSite && v.value != RecipeTypeCustom {
		return ErrInvalidRecipeType
	}
	return nil
}

func (v *recipeURL) validate(rType string) error {
	if rType == RecipeTypeSite && v.value == "" {
		return ErrInvalidRecipeURL
	}
	return nil
}

func (v *recipeMarkdown) validate(rType string) error {
	if rType == RecipeTypeCustom && v.value == "" {
		return ErrInvalidRecipeMarkdown
	}
	return nil
}

// ゲッター
func (r *Recipe) ID() string           { return r.id.value }
func (r *Recipe) DishID() string       { return r.dishID.value }
func (r *Recipe) Title() string        { return r.title.value }
func (r *Recipe) Type() string         { return r.recipeType.value }
func (r *Recipe) URL() string          { return r.url.value }
func (r *Recipe) Markdown() string     { return r.markdown.value }
func (r *Recipe) CreatedAt() time.Time { return r.createdAt }
func (r *Recipe) UpdatedAt() time.Time { return r.updatedAt }

// ドメインロジック
func (r *Recipe) IsSite() bool   { return r.recipeType.value == RecipeTypeSite }
func (r *Recipe) IsCustom() bool { return r.recipeType.value == RecipeTypeCustom }

// Update はレシピを更新する
func (r *Recipe) Update(title string, url, markdown *string) error {
	t, err := newRecipeTitle(title)
	if err != nil {
		return err
	}
	r.title = *t

	if r.IsSite() {
		if url == nil || *url == "" {
			return ErrInvalidRecipeURL
		}
		r.url = recipeURL{value: *url}
	}

	if r.IsCustom() {
		if markdown == nil || *markdown == "" {
			return ErrInvalidRecipeMarkdown
		}
		r.markdown = recipeMarkdown{value: *markdown}
	}

	r.updatedAt = time.Now()
	return nil
}

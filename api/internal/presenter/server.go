package presenter

import (
	"context"

	"recipe/api/internal/controller/auth"
	"recipe/api/internal/controller/dish"
	"recipe/api/internal/controller/recipe"
	"recipe/api/internal/controller/system"
	"recipe/api/internal/controller/user"
	"recipe/api/internal/middleware"
	"recipe/api/internal/middleware/auth0"

	"github.com/gin-gonic/gin"
)

const (
	latest        = "/v1"
	Auth0Domain   = "dev-cx-3942s.us.auth0.com"
	Auth0ClientID = "YOUR_CLIENT_ID"
)

type Server struct{}

func (s *Server) Run(ctx context.Context) error {
	// Auth0初期化
	jwks, err := auth0.FetchJWKS(Auth0Domain)
	if err != nil {
		return err
	}

	r := gin.Default()

	// CORS設定（フロントエンドからのアクセス許可）
	r.Use(middleware.CORS())

	v1 := r.Group(latest)

	// 死活監視用
	{
		systemHandler := system.NewSystemHandler()

		v1.GET("/health", systemHandler.Health)
	}

	// google認証
	{
		authHandler := auth.NewAuthHandler()

		v1.POST("/auth/google", authHandler.GoogleLogin)
	}

	authRequired := v1.Group("")
	// JWT認証Middlewareを適用
	authRequired.Use(middleware.NewAuth(jwks, Auth0Domain, Auth0ClientID))

	// user
	{
		userHandler := user.NewUserHandler()

		authRequired.GET("/user/me", userHandler.Me)
	}

	// dishes
	{
		dishHandler := dish.NewDishHandler()

		authRequired.GET("/dishes", dishHandler.List)
		authRequired.POST("/dishes", dishHandler.Create)
		authRequired.DELETE("/dishes/:dishId", dishHandler.Delete)
	}

	// recipes
	{
		recipeHandler := recipe.NewRecipeHandler()

		authRequired.GET(
			"/dishes/:dishId/recipes",
			recipeHandler.ListByDish,
		)
		authRequired.POST(
			"/dishes/:dishId/recipes",
			recipeHandler.Create,
		)
		authRequired.GET(
			"/recipes/:recipeId",
			recipeHandler.Get,
		)
		authRequired.PUT(
			"/recipes/:recipeId",
			recipeHandler.Update,
		)
		authRequired.DELETE(
			"/recipes/:recipeId",
			recipeHandler.Delete,
		)

		err := r.Run()
		if err != nil {
			return err
		}

		return nil
	}
}

func NewServer() *Server {
	return &Server{}
}

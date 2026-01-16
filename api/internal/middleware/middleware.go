package middleware

import (
	"context"
	"net/http"

	"recipe/api/internal/middleware/auth0"

	"github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

// Gin用のCORSミドルウェア
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// 初期化済みのjwtMiddlewareを受け取り、認証を行うGinミドルウェアを返す
func NewAuth(mw *jwtmiddleware.JWTMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		// jwtmiddlewareをGinに適応させる
		
		success := false
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			success = true
			c.Request = r // 更新されたコンテキストを持つリクエストをキャプチャする
		})

		mw.Handler(next).ServeHTTP(c.Writer, c.Request)

		if !success {
			// ここに来る場合、jwtmiddlewareがw.Write（拒否）を呼び出し、nextを呼び出さなかった可能性が高い
			c.Abort()
			return
		}

		// コンテキストに "user" トークンがある。これをハンドラからのアクセスのために auth0.JWTKey に伝播させる。
		userProp := mw.Options.UserProperty
		if userProp == "" {
			userProp = "user"
		}

		val := c.Request.Context().Value(userProp)
		if val != nil {
			if token, ok := val.(*jwt.Token); ok {
				// auth0.GetJWT() が見つけられるように auth0.JWTKey に設定する必要がある
				ctx := context.WithValue(c.Request.Context(), auth0.JWTKey{}, token)
				c.Request = c.Request.WithContext(ctx)
			}
		}

		c.Next()
	}
}

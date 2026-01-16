package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"recipe/api/internal/middleware/auth0"

	"github.com/gin-gonic/gin"
)

// CORSはCORS対応ミドルウェアを返す
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

// NewAuth はAuth0 JWTトークンの検証を行うミドルウェアを返す
func NewAuth(jwks *auth0.JWKS, domain, clientID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractToken(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid authorization header"})
			c.Abort()
			return
		}

		claims, err := auth0.VerifyToken(token, jwks, domain, clientID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// クレームをコンテキストに設定
		ctx := context.WithValue(c.Request.Context(), auth0.JWTKey{}, claims)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// extractToken はAuthorizationヘッダーからJWTトークンを抽出する
func extractToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format")
	}

	return parts[1], nil
}

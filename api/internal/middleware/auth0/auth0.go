package auth0

import (
	"context"
	"errors"
	"fmt"

	"github.com/form3tech-oss/jwt-go"
)

type JWTKey struct{}

// VerifyToken はJWTトークンを検証してクレームを返す
func VerifyToken(tokenString string, jwks *JWKS, domain, clientID string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 署名方式の確認
		if token.Method != jwt.SigningMethodRS256 {
			return nil, errors.New("invalid signing method")
		}

		// PEM証明書を取得
		cert, err := getPemCert(jwks, token)
		if err != nil {
			return nil, err
		}

		// RSA公開鍵をパース
		return jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	})

	if err != nil {
		return nil, fmt.Errorf("token parsing failed: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims type")
	}

	// Authorize Party (azp) の検証
	azp, ok := claims["azp"].(string)
	if !ok {
		return nil, errors.New("authorized parties are required")
	}
	if azp != clientID {
		return nil, errors.New("invalid authorized parties")
	}

	// Issuer (iss) の検証
	iss := fmt.Sprintf("https://%s/", domain)
	ok = claims.VerifyIssuer(iss, true)
	if !ok {
		return nil, errors.New("invalid issuer")
	}

	return claims, nil
}

func getPemCert(jwks *JWKS, token *jwt.Token) (string, error) {
	cert := ""

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		return "", errors.New("unable to find appropriate key")
	}

	return cert, nil
}

// GetClaims はコンテキストからJWTクレームを取得する
func GetClaims(ctx context.Context) jwt.MapClaims {
	claims, ok := ctx.Value(JWTKey{}).(jwt.MapClaims)
	if !ok {
		return nil
	}
	return claims
}

// GetUserID はコンテキストからユーザーID（sub）を取得する
// 認証済みリクエストでユーザーを識別するために使用
func GetUserID(ctx context.Context) string {
	claims := GetClaims(ctx)
	if claims == nil {
		return ""
	}
	sub, ok := claims["sub"].(string)
	if !ok {
		return ""
	}
	return sub
}

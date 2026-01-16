package auth0

// 参考: https://qiita.com/kourin1996/items/7b79d868de5c126d01d3

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JKWS向けの構造体定義
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type JWKS struct {
	Keys []JSONWebKeys `json:"keys"`
}

func FetchJWKS(auth0Domain string) (*JWKS, error) {
	resp, err := http.Get(fmt.Sprintf("https://%s/.well-known/jwks.json", auth0Domain))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jwks := &JWKS{}
	err = json.NewDecoder(resp.Body).Decode(jwks)

	return jwks, err
}

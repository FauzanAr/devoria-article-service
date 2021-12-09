package middleware

import (
	"net/http"
	"strings"

	"github.com/sangianpatrick/devoria-article-service/jwt"
)

type BearerAuth struct {
	jsonWebToken jwt.JSONWebToken
}

func NewBearerAuth(jsonWebToken jwt.JSONWebToken) RouteMiddleware {
	return &BearerAuth{jsonWebToken}
}

func ExtractToken(r *http.Request) string {
  bearToken := r.Header.Get("Authorization")
  //normally Authorization the_token_xxx
  strArr := strings.Split(bearToken, " ")
  if len(strArr) == 2 {
     return strArr[1]
  }
  return ""
}

func (ba *BearerAuth) Verify(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()
		tokenString := ExtractToken(r)

		claims, err := ba.jsonWebToken.Parse(ctx, tokenString)

		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
		}

		email, _ := claims["email"].(string)
		r.Header.Set("userEmail", email)
		
		next(w, r)
	})
}
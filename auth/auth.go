package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func AuthCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorized := ValidToken(r.Header.Get("Authorization"))
		if authorized || ExcludedRoute(r.URL.Path) {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized"))
			return
		}
	})
}

func ExcludedRoute(url string) bool {
	switch url {
	case "/api/v1/checktoken":
		return true
	case "/api/v1/users":
		return true
	case "/api/v1/login":
		return true
	default:
		return false
	}
}

func ValidToken(tokenString string) bool {
	secret, ok := os.LookupEnv("SECRET")
	if !ok {
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		fmt.Println("Error validating token or invalid token:", err)
		return false
	}
	return true
}

func CreateToken(email string) string {
	secret, ok := os.LookupEnv("SECRET")
	if !ok {
		return ""
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Error creating token:", err)
	}
	return tokenString
}

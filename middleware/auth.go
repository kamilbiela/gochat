package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/justinas/alice"
	"github.com/kamilbiela/gochat/lib"
	"net/http"
	"time"
)

type authErrorResponse struct {
	Message string
}

func writeAuthError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusForbidden)
	w.Header().Set("Content-Type", "application/json")

	j, _ := json.Marshal(authErrorResponse{Message: msg})
	fmt.Fprint(w, string(j))
}

func Auth(auth *lib.AuthService) alice.Constructor {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()

			tokensStr, ok := q["token"]

			if !ok {
				writeAuthError(w, "Missing token")
				return
			}

			tokenStr := string(tokensStr[0])

			tokenObj := auth.GetToken(tokenStr)

			if tokenObj == nil {
				writeAuthError(w, "Invalid token")
				return
			}

			if time.Now().After(tokenObj.ExpireAt) {
				writeAuthError(w, "Token expired")
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
}

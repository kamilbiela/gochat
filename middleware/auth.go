package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/kamilbiela/gochat/lib"
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
			handler.ServeHTTP(w, r)
			return
			q := r.URL.Query()
			var tokenStr string

			tokensStr, ok := q["token"]

			if !ok {
				// try to find token in url
				vars := mux.Vars(r)
				tokenStr, ok = vars["token"]

				log.Println("/=========a\\")
				log.Println(vars)
				log.Println(tokenStr, ok)
				log.Println("\\=========a/")
				if !ok {
					writeAuthError(w, "Missing token")
					return
				}
			} else {
				tokenStr = string(tokensStr[0])
			}

			tokenObj := auth.GetToken(tokenStr)

			if tokenObj == nil {
				writeAuthError(w, "Invalid token")
				return
			}

			if !tokenObj.IsExpired() {
				writeAuthError(w, "Token expired")
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
}

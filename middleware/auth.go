package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/kamilbiela/gochat/lib"
	"log"
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
			var tokenStr string

			tokensStr, ok := q["token"]

			if !ok {
				// try to find token in url
				vars := mux.Vars(r)
				tokenStr, ok = vars["token"]

				log.Println(vars)
				log.Println(tokenStr, ok)
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

			if time.Now().After(tokenObj.ExpireAt) {
				writeAuthError(w, "Token expired")
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
}

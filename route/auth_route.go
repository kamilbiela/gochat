package route

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/kamilbiela/gochat/lib"
	"net/http"
)

type AuthRequest struct {
	Username string
	Password string
}

type AuthResponse struct {
	Token *lib.Token
}

func (ar *AuthRequest) isValid(auth *lib.AuthService) bool {
	return !govalidator.IsNull(ar.Username) && !govalidator.IsNull(ar.Password) &&
		auth.IsValid(ar.Username, ar.Password)
}

func AuthRoute(auth *lib.AuthService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authRequest := new(AuthRequest)

		if err := json.NewDecoder(r.Body).Decode(authRequest); err != nil {
			lib.WriteError(w, 400, err.Error())
			return
		}

		if !authRequest.isValid(auth) {
			lib.WriteError(w, 400, "Bad credentials")
			return
		}

		Token := auth.GenerateToken(authRequest.Username)

		response := AuthResponse{
			Token: &Token,
		}

		lib.WriteSuccess(w, response)
	})
}

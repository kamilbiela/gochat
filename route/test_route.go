package route

import (
	"fmt"
	"net/http"
)

func TestRoute(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, message)
	})
}

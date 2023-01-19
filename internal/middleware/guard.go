package middleware

import (
	"errors"
	"github.com/thanh-vt/splash-inventory-service/internal/constant"
	"net/http"
)

func HandleGuard(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		currentUser := r.Context().Value(constant.UID)
		if currentUser == nil {
			Unauthorized(w, r, errors.New("not logged in yet"))
			return
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)

}

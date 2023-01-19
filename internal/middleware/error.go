package middleware

import (
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"net/http"
	"runtime/debug"
	"time"
)

func Unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	render.Status(r, http.StatusUnauthorized)
	render.JSON(w, r, renderError(r, err))
}

func BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, renderError(r, err))
}

func NotFound(w http.ResponseWriter, r *http.Request, err error) {
	render.Status(r, http.StatusNotFound)
	render.JSON(w, r, renderError(r, err))
}

func InternalServer(w http.ResponseWriter, r *http.Request, err error) {
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, renderError(r, err))
}

func renderError(r *http.Request, err error) map[string]interface{} {
	return map[string]interface{}{
		"message":   err.Error(),
		"timestamp": time.Now(),
		"path":      r.RequestURI,
	}
}

func HandleError(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				if rvr == http.ErrAbortHandler {
					// we don't recover http.ErrAbortHandler so the response
					// to the client is aborted, this should not be logged
					panic(rvr)
				}

				logEntry := middleware.GetLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(rvr, debug.Stack())
				} else {
					middleware.PrintPrettyStack(rvr)
				}
				w.Header().Set("content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				errRes, err := json.Marshal(map[string]interface{}{
					"message":   rvr,
					"timestamp": time.Now(),
					"path":      r.RequestURI,
				})
				_, err = w.Write(errRes)
				if err != nil {
					return
				}
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)

}

package handler

import (
	"net/http"
	"study/auth"
)

func AuthMiddleware(j *auth.JWTer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req, err := j.FillContext(r)
			if err != nil {
				ResposndJSON(r.Context(), w, ErrResponse{
					Message: "not find auth info",
				}, http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, req)
		})
	}
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !auth.IsAdmin(r.Context()) {
			ResposndJSON(r.Context(), w, ErrResponse{
				Message: "not admin",
			}, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

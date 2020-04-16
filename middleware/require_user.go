package middleware

import (
	"fmt"	
	"net/http"
	"github.com/igorvinnicius/lenslocked-go-web/models"
	"github.com/igorvinnicius/lenslocked-go-web/context"
)

type RequireUser struct {
	models.UserService
}

func (mw *RequireUser) Apply(next http.Handler) http.HandlerFunc {
	return mw.ApplyFn(next.ServeHTTP)
}

func (mw *RequireUser) ApplyFn(next http.HandlerFunc) http.HandlerFunc {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		
		cookie, err := r.Cookie("remember_token")

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		user, err := mw.UserService.ByRemember(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		
		ctx := r.Context()
		ctx = context.WithUser(ctx, user)

		r = r.WithContext(ctx)

		fmt.Println(user)

		next(w, r)
	})
}
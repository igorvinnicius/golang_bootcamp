package middleware

import (
	"fmt"
	"github.com/igorvinnicius/lenslocked-go-web/context"
	"github.com/igorvinnicius/lenslocked-go-web/models"
	"net/http"
)

type User struct {
	models.UserService
}

type RequireUser struct {
	User
}

func (mw *User) Apply(next http.Handler) http.HandlerFunc {
	return mw.ApplyFn(next.ServeHTTP)
}

func (mw *RequireUser) ApplyFn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := context.User(r.Context())
		if user == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next(w, r)
	})
}


func (mw *User) ApplyFn(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		cookie, err := r.Cookie("remember_token")
		if err != nil {
			next(w, r)
			fmt.Println(err)
			return
		}

		user, err := mw.UserService.ByRemember(cookie.Value)
		if err != nil {
			next(w, r)
			fmt.Println(err)
			return
		}

		ctx := r.Context()
		ctx = context.WithUser(ctx, user)
		r = r.WithContext(ctx)
		next(w, r)
	})
}
package auth

import (
	"alshashiguchi/quiz_gem/graph/model"
	users "alshashiguchi/quiz_gem/models/users"
	"context"
	"net/http"

	jwt "alshashiguchi/quiz_gem/core/security"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	user string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			user, err := users.GetUserByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctxUser := users.User{Username: user.Username, Access: user.Access, Email: user.Email, ID: user.ID, Name: user.Name, Situation: user.Situation}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &ctxUser)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context, access []model.Access) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)

	if !contains(access, raw.Access) {
		return nil
	}

	return raw
}
func contains(listAccess []model.Access, userAccess model.Access) bool {
	for _, a := range listAccess {
		if a == userAccess {
			return true
		}
	}
	return false
}

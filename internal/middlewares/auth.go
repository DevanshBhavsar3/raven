package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkHTTP "github.com/clerk/clerk-sdk-go/v2/http"
)

type UserIDContextKey struct{}
type UserRoleContextKey struct{}
type UserPermissionsContextKey struct{}

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (auth *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return clerkHTTP.WithHeaderAuthorization(
		clerkHTTP.AuthorizationFailureHandler(
			http.HandlerFunc(auth.authorizationFailure),
		),
	)(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			claims, ok := clerk.SessionClaimsFromContext(ctx)
			if !ok || claims == nil {
				auth.authorizationFailure(w, r)
				return
			}

			ctx = context.WithValue(ctx, UserIDContextKey{}, claims.Subject)
			ctx = context.WithValue(ctx, UserRoleContextKey{}, claims.ActiveOrganizationRole)
			ctx = context.WithValue(ctx, UserPermissionsContextKey{}, claims.ActiveOrganizationPermissions)

			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		}),
	)
}

func (auth *AuthMiddleware) authorizationFailure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	w.Write([]byte("unauthorized"))
}

func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(UserIDContextKey{}).(string)
	return userID, ok
}

func GetUserRoleFromContext(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(UserRoleContextKey{}).(string)
	return role, ok
}

func GetUserPermissionsFromContext(ctx context.Context) ([]string, bool) {
	permissions, ok := ctx.Value(UserPermissionsContextKey{}).([]string)
	return permissions, ok
}

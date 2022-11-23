package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"samsul96maarif/github.com/go-api-app/config"
	"samsul96maarif/github.com/go-api-app/lib"
	"samsul96maarif/github.com/go-api-app/lib/logger"
	"samsul96maarif/github.com/go-api-app/request"

	"github.com/felixge/httpsnoop"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func (handler *Handler) PublicMiddleware(next http.Handler) http.Handler {
	return handler.StandardMiddleware(next)
}

func (handler *Handler) ProtectedMiddleware(next http.Handler) http.Handler {
	return handler.StandardMiddleware(handler.AuthMiddleware(next))
}

func (handler *Handler) ProtectedByRolesMiddleware(next http.Handler, roles []int) http.Handler {
	return handler.StandardMiddleware(handler.AuthMiddleware(handler.ValidateRolesMiddleware(next, roles)))
}

func (handler *Handler) PanicMiddlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				logger.Error(r.Context(), "panic occured", map[string]interface{}{
					"error": rec,
				})
				writeError(w, lib.ErrorInternalServer)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (handler *Handler) StandardMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		ctx := context.WithValue(r.Context(), "X-Request-ID", uuid.NewString())

		m := httpsnoop.CaptureMetrics(handler.PanicMiddlewares(next), w, r.WithContext(ctx))

		logger.Info(ctx, "http api request", map[string]interface{}{
			"method":   r.Method,
			"path":     r.URL,
			"status":   m.Code,
			"duration": m.Duration.Milliseconds(),
		})
	})
}

func (handler *Handler) ValidateRolesMiddleware(next http.Handler, allowedRoleIDs []int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		splitToken := strings.Split(authHeader, "Bearer ")

		if len(splitToken) < 2 {
			writeError(w, lib.ErrorUnauthorized)
			return
		}

		claim, ok := r.Context().Value("User").(lib.MyClaim)
		if !ok {
			writeError(w, lib.ErrorUnauthorized)
			return
		}

		fmt.Printf(" claim %+v \n", claim)

		var isUserRole bool
		var userRoleFounded uint
		for _, role := range allowedRoleIDs {
			for _, userRole := range claim.Roles {
				if role == userRole {
					isUserRole = true
					userRoleFounded = uint(role)
					break
				}
			}
			if isUserRole {
				break
			}
		}

		if !isUserRole {
			writeError(w, lib.ErrorForbidden)
			return
		}

		user, err := handler.BE.Usecase.FindUser(r.Context(), request.FindUserRequest{Email: claim.Email})
		if err != nil || user.Id == 0 {
			writeError(w, lib.ErrorUnauthorized)
			return
		}

		userRole, err := handler.BE.Usecase.FindUserRole(r.Context(), request.FindUserRoleRequest{
			UserId: user.Id,
			RoleId: userRoleFounded,
		})
		if err != nil || userRole.UserId == 0 {
			writeError(w, lib.ErrorForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (handler *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenStr := strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != config.JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}
			return config.GetSignatureKey(), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		email, _ := claims["email"].(string)
		claimRoles, _ := claims["roles"].([]interface{})
		userId, _ := claims["user_id"].(uint)
		var roles []int
		for _, val := range claimRoles {
			var valInt int
			switch i := val.(type) {
			case float64:
				valInt = int(i)
				break
			case float32:
				valInt = int(i)
				break
				//default:
				//	return math.NaN(), errors.New("getFloat: unknown value is of incompatible type")
			}
			roles = append(roles, valInt)
		}

		customClaim := lib.MyClaim{
			Roles:  roles,
			Email:  email,
			UserId: userId,
		}

		ctx := context.WithValue(context.Background(), "User", customClaim)
		r = r.WithContext(ctx)

		logger.Info(ctx, "claim", map[string]interface{}{
			"claims": claims,
		})

		next.ServeHTTP(w, r)
	})
}

package middlewares

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/dollarkillerx/common/pkg/jwt"
	"github.com/dollarkillerx/fireworks/internal/config"
	"github.com/dollarkillerx/fireworks/utils"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// contextKey ...
// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
type contextKey string

func (c contextKey) String() string {
	return string(c)
}

// contextKey ...
const (
	tokenCtxKey             contextKey = "Token"
	requestReceivedAtCtxKey contextKey = "ReqReceivedAt"
)

// Context  get user from jwt and put user into ctx
func Context() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := uuid.New().String()
			ctx := utils.SetContextRequestID(r.Context(), reqID)
			ctx = context.WithValue(ctx, requestReceivedAtCtxKey, time.Now())
			// token
			tokenString, _ := getTokenFromHeader(r.Header)
			if tokenString != "" {
				ctx = context.WithValue(ctx, tokenCtxKey, tokenString)
			}

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// GetUserIDFromCtx ...
func GetUserIDFromCtx(ctx context.Context) (string, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return "", err
	}
	if userID != nil {
		return *userID, nil
	}

	tokenString, _ := GetTokenFromCtx(ctx)
	userInfo, err := jwt.GetInfoFromAccessToken(tokenString, config.GetJWTConfig().SecretKey)
	if err != nil {
		return "", errors.Errorf("TOKEN 错误")
	}

	s, ex := userInfo.Info["UserID"]
	if !ex {
		return "", errors.New("TOKEN 错误")
	}

	return s, nil
}

// GetUserIDFromContext ...
func GetUserIDFromContext(ctx context.Context) (*string, error) {
	value := ctx.Value(userIDCtxKey)
	if value == nil {
		return nil, nil
	}
	userID, ok := value.(string)
	if !ok {
		return nil, errors.New("UserID in context is not string")
	}
	return &userID, nil
}

func SetContextUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDCtxKey, userID)
}

// GetTokenFromCtx ...
func GetTokenFromCtx(ctx context.Context) (string, error) {
	tokenString, ok := ctx.Value(tokenCtxKey).(string)
	if !ok {
		return "", errors.New("token not found in context")
	}
	return tokenString, nil
}

// 从请求中获取 access token
func getTokenFromHeader(header http.Header) (string, error) {
	authorization := header.Get("Authorization")
	if authorization == "" {
		return "", errors.New("Invalid user. No Authorization found")
	}
	tokenString := strings.TrimSpace(strings.Replace(authorization, "Bearer", "", 1))
	if tokenString == "" {
		return "", errors.New("Invalid user. No Bearer token string found")
	}
	return tokenString, nil
}

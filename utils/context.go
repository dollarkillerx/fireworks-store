package utils

import (
	"context"

	"github.com/pkg/errors"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

const (
	requestID contextKey = "RequestID"
)

func SetContextRequestID(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, requestID, reqID)
}

func GetContextRequestID(ctx context.Context) string {
	value := ctx.Value(requestID)
	if value == nil {
		return ""
	}
	reqID, ok := value.(string)
	if !ok {
		return ""
	}
	return reqID
}

func GetReqIDFromContext(ctx context.Context) (*string, error) {
	value := ctx.Value(requestID)
	if value == nil {
		return nil, nil
	}
	reqID, ok := value.(string)
	if !ok {
		return nil, errors.New("ReqID in context is not string")
	}
	return &reqID, nil
}

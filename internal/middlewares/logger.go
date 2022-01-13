package middlewares

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/99designs/gqlgen/graphql"
	utils2 "github.com/dollarkillerx/fireworks/pkg/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ErrorPresenter ...
func ErrorPresenter(ctx context.Context, e error) *gqlerror.Error {
	utils2.Logger.Errorf("[ReqID]:%s\n[Message]:%+v", utils2.GetContextRequestID(ctx), e)

	return graphql.DefaultErrorPresenter(ctx, e)
}

// RecoverFunc ...
func RecoverFunc(ctx context.Context, err interface{}) (userMessage error) {
	utils2.Logger.Errorf("[ReqID]:%s\n[Message]:%+v\n[Panic]:%s", utils2.GetContextRequestID(ctx), err, string(debug.Stack()))
	return fmt.Errorf("panic: %+v", err)
}

// ResponseLogger ...
func ResponseLogger(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	res := next(ctx)
	if res == nil {
		return res
	}
	if res.Errors != nil {
		return res
	}
	if res.Data == nil {
		return res
	}
	receivedAt, _ := ctx.Value(RequestReceivedAtCtxKey).(time.Time)
	utils2.Logger.Tracef("[ReqID]:%s\n[Duration]:%dms\n[Response]:%s",
		utils2.GetContextRequestID(ctx),
		time.Since(receivedAt).Milliseconds(),
		string(res.Data),
	)
	return res
}

// RequestLogger ...
type RequestLogger struct{}

// ExtensionName ...
func (r *RequestLogger) ExtensionName() string {
	return "RequestLogger"
}

// Validate ...
func (r *RequestLogger) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

// MutateOperationParameters ...
func (r *RequestLogger) MutateOperationParameters(ctx context.Context, request *graphql.RawParams) *gqlerror.Error {
	if request == nil {
		return nil
	}
	if request.OperationName == "IntrospectionQuery" {
		return nil
	}
	content := fmt.Sprintf("[ReqID]:%s\n[Query]:%s\n[OperationName]:%s",
		utils2.GetContextRequestID(ctx),
		request.Query,
		request.OperationName,
	)
	if len(request.Variables) > 0 {
		content = content + fmt.Sprintf("\n[Variables]:%+v", request.Variables)
	}
	if len(request.Extensions) > 0 {
		content = content + fmt.Sprintf("\n[Extensions]:%+v", request.Extensions)
	}
	utils2.Logger.Traceln(content)
	return nil
}

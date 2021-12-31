package middlewares

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dollarkillerx/fireworks/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ErrorPresenter ...
func ErrorPresenter(ctx context.Context, e error) *gqlerror.Error {
	utils.Logger.Errorf("[ReqID]:%s\n[Message]:%+v", utils.GetContextRequestID(ctx), e)

	return graphql.DefaultErrorPresenter(ctx, e)
}

// RecoverFunc ...
func RecoverFunc(ctx context.Context, err interface{}) (userMessage error) {
	utils.Logger.Errorf("[ReqID]:%s\n[Message]:%+v\n[Panic]:%s", utils.GetContextRequestID(ctx), err, string(debug.Stack()))
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
	receivedAt, _ := ctx.Value(requestReceivedAtCtxKey).(time.Time)
	utils.Logger.Tracef("[ReqID]:%s\n[Duration]:%dms\n[Response]:%s",
		utils.GetContextRequestID(ctx),
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
		utils.GetContextRequestID(ctx),
		request.Query,
		request.OperationName,
	)
	if len(request.Variables) > 0 {
		content = content + fmt.Sprintf("\n[Variables]:%+v", request.Variables)
	}
	if len(request.Extensions) > 0 {
		content = content + fmt.Sprintf("\n[Extensions]:%+v", request.Extensions)
	}
	utils.Logger.Traceln(content)
	return nil
}

package middlewares

type ctxKey string

func (c ctxKey) String() string {
	return string(c)
}

const (
	userIDCtxKey      ctxKey = "UserID"
	userNameCtxKey    ctxKey = "UserName"
	reqIDCtxKey       ctxKey = "ReqID"
	serviceNameCtxKey ctxKey = "ServiceName"
	tagCtxKey         ctxKey = "Tag"
)

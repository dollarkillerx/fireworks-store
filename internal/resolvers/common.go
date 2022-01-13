package resolvers

import (
	"context"
	"sync"
	"time"

	"github.com/dollarkillerx/fireworks/internal/generated"
	"github.com/dollarkillerx/fireworks/internal/storage"
	"github.com/dollarkillerx/fireworks/internal/storage/basis"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewResolver(storage *basis.Storage) *Resolver {
	r := &Resolver{
		storage: storage,
		subscriptions: make(map[string][]struct {
			ID string
		}),
	}
	return r
}

type Resolver struct {
	sync.Mutex
	storage       storage.Interface
	subscriptions map[string][]struct {
		ID string
	}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver } // MutationResolver

func (m *mutationResolver) CreateCommodity(ctx context.Context, input generated.CreateCommodity) (bool, error) {
	panic("implement me")
}

func (m *mutationResolver) UpdateProduct(ctx context.Context, input generated.UpdateCommodity) (bool, error) {
	panic("implement me")
}

func (m *mutationResolver) IbRegistry(ctx context.Context, input generated.IbRegistry) (*generated.AuthPayload, error) {
	panic("implement me")
}

func (m *mutationResolver) WithdrawalCommission(ctx context.Context, alipayID string, captcha string, captchaToken string) (bool, error) {
	panic("implement me")
}

func (m *mutationResolver) CaptchaImg(ctx context.Context) (*generated.Captcha, error) {
	panic("implement me")
}

func (m *mutationResolver) CaptchaPhone(ctx context.Context, phone string) (*generated.Captcha, error) {
	panic("implement me")
}

func (m *mutationResolver) PlaceOrder(ctx context.Context, commodityID []string, shippingAddressID string) (*generated.PlaceOrder, error) {
	panic("implement me")
}

func (m *mutationResolver) CreateShippingAddress(ctx context.Context, input generated.CreateShippingAddress) (*bool, error) {
	panic("implement me")
}

func (m *mutationResolver) RemoveShippingAddress(ctx context.Context, id string) (*bool, error) {
	panic("implement me")
}

func (m *mutationResolver) HealthCheck(ctx context.Context) (string, error) {
	return "ok", nil
}

func (m *mutationResolver) RegistrationInvitationCode(ctx context.Context, invitationCode string) (*bool, error) {
	//TODO implement me
	panic("implement me")
}

type queryResolver struct{ *Resolver } // QueryResolver

func (q *queryResolver) AreaList(ctx context.Context, latitude int, longitude int) ([]generated.AreaList, error) {
	panic("implement me")
}

func (q *queryResolver) BdLogin(ctx context.Context, account string, password string, captcha string, captchaToken string) (*generated.AuthPayload, error) {
	panic("implement me")
}

func (q *queryResolver) BdCommodity(ctx context.Context, keyword *string, pagination generated.Pagination) (*generated.CommodityList, error) {
	panic("implement me")
}

func (q *queryResolver) BdProductDetails(ctx context.Context, commodityID string) (*generated.Commodity, error) {
	panic("implement me")
}

func (q *queryResolver) IbLogin(ctx context.Context, account string, password string, captcha string, captchaToken string) (*generated.AuthPayload, error) {
	panic("implement me")
}

func (q *queryResolver) RebateMerchandise(ctx context.Context, keyword *string, pagination generated.Pagination) (*generated.CommodityList, error) {
	panic("implement me")
}

func (q *queryResolver) MyRebate(ctx context.Context) (*generated.MyRebate, error) {
	panic("implement me")
}

func (q *queryResolver) MyIb(ctx context.Context) (*generated.MyIb, error) {
	panic("implement me")
}

func (q *queryResolver) RebateFlow(ctx context.Context, startTime int, endTime int, pagination generated.Pagination) (*generated.RebateFlow, error) {
	panic("implement me")
}

func (q *queryResolver) UserFlow(ctx context.Context, startTime int, endTime int, pagination generated.Pagination) (*generated.UserFlow, error) {
	panic("implement me")
}

func (q *queryResolver) CommodityList(ctx context.Context, keyword *string, shopID string, bought bool, pagination generated.Pagination) (*generated.CommodityList, error) {
	panic("implement me")
}

func (q *queryResolver) Order(ctx context.Context, id string) (*generated.Order, error) {
	panic("implement me")
}

func (q *queryResolver) HistoryOrder(ctx context.Context, pagination generated.Pagination) (*generated.HistoryOrder, error) {
	panic("implement me")
}

func (q *queryResolver) ShippingAddress(ctx context.Context) (*generated.ShippingAddressList, error) {
	panic("implement me")
}

func (q *queryResolver) HealthCheck(ctx context.Context) (string, error) {
	return "ok", nil
}

func (q *queryResolver) Now(ctx context.Context) (*timestamppb.Timestamp, error) {
	return &timestamppb.Timestamp{Seconds: time.Now().Unix()}, nil
}

func (q *queryResolver) User(ctx context.Context) (*generated.UserInformation, error) {
	return &generated.UserInformation{
		UserID: "xxx",
	}, nil
}

type subscriptionResolver struct{ *Resolver }

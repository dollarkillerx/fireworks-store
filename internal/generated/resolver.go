package generated

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Resolver struct{}

func (r *mutationResolver) HealthCheck(ctx context.Context) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateCommodity(ctx context.Context, input CreateCommodity) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, input UpdateCommodity) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) IbRegistry(ctx context.Context, input IbRegistry) (*AuthPayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) WithdrawalCommission(ctx context.Context, alipayID string, captcha string, captchaToken string) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) CaptchaImg(ctx context.Context) (*Captcha, error) {
	panic("not implemented")
}

func (r *mutationResolver) CaptchaPhone(ctx context.Context, phone string) (*Captcha, error) {
	panic("not implemented")
}

func (r *mutationResolver) PlaceOrder(ctx context.Context, commodityID []string, shippingAddressID string) (*PlaceOrder, error) {
	panic("not implemented")
}

func (r *mutationResolver) BasicRegistrationInformation(ctx context.Context, input *BasicInfo) (*bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) RegistrationInvitationCode(ctx context.Context, invitationCode string) (*bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateShippingAddress(ctx context.Context, input CreateShippingAddress) (*bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) RemoveShippingAddress(ctx context.Context, id string) (*bool, error) {
	panic("not implemented")
}

func (r *queryResolver) HealthCheck(ctx context.Context) (string, error) {
	panic("not implemented")
}

func (r *queryResolver) Now(ctx context.Context) (*timestamppb.Timestamp, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context) (*UserInformation, error) {
	panic("not implemented")
}

func (r *queryResolver) UserLogin(ctx context.Context, token string, latitude int, longitude int) (*AuthPayload, error) {
	panic("not implemented")
}

func (r *queryResolver) AreaList(ctx context.Context, latitude int, longitude int) ([]AreaList, error) {
	panic("not implemented")
}

func (r *queryResolver) BdLogin(ctx context.Context, account string, password string, captcha string, captchaToken string) (*AuthPayload, error) {
	panic("not implemented")
}

func (r *queryResolver) BdCommodity(ctx context.Context, keyword *string, pagination Pagination) (*CommodityList, error) {
	panic("not implemented")
}

func (r *queryResolver) BdProductDetails(ctx context.Context, commodityID string) (*Commodity, error) {
	panic("not implemented")
}

func (r *queryResolver) IbLogin(ctx context.Context, account string, password string, captcha string, captchaToken string) (*AuthPayload, error) {
	panic("not implemented")
}

func (r *queryResolver) RebateMerchandise(ctx context.Context, keyword *string, pagination Pagination) (*CommodityList, error) {
	panic("not implemented")
}

func (r *queryResolver) MyRebate(ctx context.Context) (*MyRebate, error) {
	panic("not implemented")
}

func (r *queryResolver) MyIb(ctx context.Context) (*MyIb, error) {
	panic("not implemented")
}

func (r *queryResolver) RebateFlow(ctx context.Context, startTime int, endTime int, pagination Pagination) (*RebateFlow, error) {
	panic("not implemented")
}

func (r *queryResolver) UserFlow(ctx context.Context, startTime int, endTime int, pagination Pagination) (*UserFlow, error) {
	panic("not implemented")
}

func (r *queryResolver) ShopList(ctx context.Context, areaID string, sales bool, pagination Pagination) (*ShopList, error) {
	panic("not implemented")
}

func (r *queryResolver) SearchShop(ctx context.Context, keyword string) (*ShopList, error) {
	panic("not implemented")
}

func (r *queryResolver) CommodityList(ctx context.Context, keyword *string, shopID string, bought bool, pagination Pagination) (*CommodityList, error) {
	panic("not implemented")
}

func (r *queryResolver) Order(ctx context.Context, id string) (*Order, error) {
	panic("not implemented")
}

func (r *queryResolver) HistoryOrder(ctx context.Context, pagination Pagination) (*HistoryOrder, error) {
	panic("not implemented")
}

func (r *queryResolver) ShippingAddress(ctx context.Context) (*ShippingAddressList, error) {
	panic("not implemented")
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

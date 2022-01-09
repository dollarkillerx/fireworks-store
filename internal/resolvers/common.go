package resolvers

import (
	"context"
	"sync"
	"time"

	"github.com/dollarkillerx/fireworks/internal/generated"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewResolver() *Resolver {

	r := &Resolver{
		subscriptions: make(map[string][]struct {
			ID string
		}),
	}
	return r
}

type Resolver struct {
	sync.Mutex
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

func (m *mutationResolver) HealthCheck(ctx context.Context) (string, error) {
	return "ok", nil
}

func (m *mutationResolver) BasicRegistrationInformation(ctx context.Context, input *generated.BasicInfo) (*bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mutationResolver) RegistrationInvitationCode(ctx context.Context, invitationCode string) (*bool, error) {
	//TODO implement me
	panic("implement me")
}

type queryResolver struct{ *Resolver } // QueryResolver

func (q *queryResolver) HealthCheck(ctx context.Context) (string, error) {
	return "ok", nil
}

func (q *queryResolver) UserLogin(ctx context.Context, token string, latitude int, longitude int) (*generated.AuthPayload, error) {
	//TODO implement me
	panic("implement me")
}

func (q *queryResolver) LocationList(ctx context.Context, latitude int, longitude int) ([]generated.LocationList, error) {
	//TODO implement me
	panic("implement me")
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

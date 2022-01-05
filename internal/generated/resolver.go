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

func (r *mutationResolver) BasicRegistrationInformation(ctx context.Context, input *BasicInfo) (*bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) RegistrationInvitationCode(ctx context.Context, invitationCode string) (*bool, error) {
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

func (r *queryResolver) LocationList(ctx context.Context, latitude int, longitude int) ([]LocationList, error) {
	panic("not implemented")
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

package generated

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Resolver struct{}

func (r *mutationResolver) Healthcheck(ctx context.Context) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateUserAccount(ctx context.Context, email string, password string) (*AuthPayload, error) {
	panic("not implemented")
}

func (r *queryResolver) Healthcheck(ctx context.Context) (string, error) {
	panic("not implemented")
}

func (r *queryResolver) Now(ctx context.Context) (*timestamppb.Timestamp, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context) (*UserInformation, error) {
	panic("not implemented")
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

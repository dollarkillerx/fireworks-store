package resolvers

import (
	"context"
	"github.com/dollarkillerx/fireworks/pkg/utils"

	"github.com/dollarkillerx/fireworks/internal/config"
	"github.com/dollarkillerx/fireworks/internal/generated"
)

func (m *mutationResolver) CreateUserAccount(ctx context.Context, email string, password string) (*generated.AuthPayload, error) {
	token, err := utils.CreateAccessToken(email, config.GetJWTConfig().SecretKey)
	if err != nil {
		return nil, err
	}

	return &generated.AuthPayload{AccessTokenString: token}, nil
}

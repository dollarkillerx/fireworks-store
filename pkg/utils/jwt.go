package utils

import (
	"github.com/dollarkillerx/common/pkg/jwt"
)

// CreateAccessToken ...
func CreateAccessToken(ID string, secretKey string) (string, error) {
	info := jwt.Info{
		map[string]string{
			"UserID": ID,
		},
	}
	return jwt.CreateAccessToken(info, secretKey)
}

// CreateAccessTokenByMap ...
func CreateAccessTokenByMap(m map[string]string, secretKey string) (string, error) {
	info := jwt.Info{
		Info: m,
	}
	return jwt.CreateAccessToken(info, secretKey)
}

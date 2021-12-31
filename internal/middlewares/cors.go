package middlewares

import (
	"github.com/dollarkillerx/fireworks/internal/config"
	"github.com/rs/cors"

	"net/http"
)

// Cors ...
func Cors() func(http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   config.GetCORSAllowedOrigins(),
		AllowedHeaders:   []string{"Authorization", "Content-Type", "isInternalServer", "OAuthorization", "TeamID"},
		AllowCredentials: true,
	}).Handler
}

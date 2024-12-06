//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yaninyzwitty/graphql-cocroach-go/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB            *pgxpool.Pool
	SocialService *service.SocialService
}

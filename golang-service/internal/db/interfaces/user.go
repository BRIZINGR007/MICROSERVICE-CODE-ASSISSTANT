package interfcaes

import (
	"context"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/repositories"
)

type UserRepositoryInterface interface {
	AddUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UserIndexes(ctx context.Context) error
}

var _ UserRepositoryInterface = (*repositories.UserRepository)(nil)

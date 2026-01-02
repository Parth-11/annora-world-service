package worldrepo

import (
	"context"

	"github.com/Parth-11/annora-world-service/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateWorld(ctx context.Context, world *models.World, creatorId uuid.UUID) error
	GetWorldById(ctx context.Context, id uuid.UUID) (*models.World, error)
	ListWorlds(ctx context.Context, limit, offset int) ([]*models.World, error)
}

type WorldRepo struct {
	DB *pgxpool.Pool
}

func NewWorldRepo(db *pgxpool.Pool) *WorldRepo {
	return &WorldRepo{
		DB: db,
	}
}

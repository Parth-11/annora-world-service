package worldrepo

import (
	"context"

	"github.com/Parth-11/annora-world-service/internal/models"
	"github.com/google/uuid"
)

func (repo *WorldRepo) GetWorldById(ctx context.Context, id uuid.UUID) (*models.World, error) {
	var world models.World

	query := `
		SELECT id,name,description, creator_id,
		created_at,updated_at,member_count,cover_img,
		visibility,status
		FROM worlds
		WHERE id = $1 AND status = 'active'
	`

	err := repo.DB.QueryRow(ctx, query, id).Scan(
		&world.ID,
		&world.Name,
		&world.Description,
		&world.CreatorID,
		&world.CreatedAt,
		&world.UpdatedAt,
		&world.MemberCount,
		&world.CoverImg,
		&world.Visibility,
		&world.Status,
	)
	if err != nil {
		return nil, err
	}

	return &world, nil
}

func (repo *WorldRepo) ListWorlds(ctx context.Context, limit, offset int) ([]*models.World, error) {
	query := `
		SELECT id,name,description, creator_id,
		created_at,updated_at,member_count,cover_img,
		visibility
		FROM worlds
		WHERE status = 'active'
		LIMIT $1 OFFSET $2
	`

	rows, err := repo.DB.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	var worlds []*models.World
	for rows.Next() {
		var world models.World
		if err := rows.Scan(
			&world.ID,
			&world.Name,
			&world.Description,
			&world.CreatorID,
			&world.CreatedAt,
			&world.UpdatedAt,
			&world.MemberCount,
			&world.CoverImg,
			&world.Visibility,
		); err != nil {
			return nil, err
		}

		world.Status = models.WorldStatusActive
		worlds = append(worlds, &world)
	}

	return worlds, rows.Err()
}

package worldrepo

import (
	"context"
	"time"

	"github.com/Parth-11/annora-world-service/internal/models"
	"github.com/google/uuid"
)

func (repo *WorldRepo) CreateWorld(ctx context.Context, world *models.World, creatorId uuid.UUID) error {
	tx, err := repo.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO worlds (
			id,name,description,creator_id,cover_img,visibility,
			status,member_count
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,1
		)
		RETURNING created_at, updated_at
	`

	err = tx.QueryRow(ctx, query,
		world.ID,
		world.Name,
		world.Description,
		world.CreatorID,
		world.CoverImg,
		world.Visibility,
		world.Status,
	).Scan(&world.CreatedAt, &world.UpdatedAt)
	if err != nil {
		return err
	}

	membershipQuery := `
		INSERT INTO world_memberships (
			world_id,user_id,role,reputation,joined_at
		) VALUES (
			$1,$2,'owner',100,$3
		)
	`

	_, err = tx.Exec(ctx, membershipQuery, world.ID, world.CreatorID, time.Now())
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

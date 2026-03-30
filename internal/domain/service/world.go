package worldservice

import worldrepo "github.com/Parth-11/annora-world-service/internal/repository/world"

type Service struct {
	WorldRepo *worldrepo.WorldRepo
}

func NewService(world_repo *worldrepo.WorldRepo) *Service {
	return &Service{}
}

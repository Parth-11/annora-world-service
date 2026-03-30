package worldhandler

import worldservice "github.com/Parth-11/annora-world-service/internal/domain/service"

type Handler struct {
	Service *worldservice.Service
}

func NewHandler(service *worldservice.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

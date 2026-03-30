package https

import (
	"net/http"

	worldhandler "github.com/Parth-11/annora-world-service/internal/api/http/world"
	"github.com/go-chi/chi/v5"
)

func NewRouter(world_handler *worldhandler.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from world service"))
	})

	return r
}

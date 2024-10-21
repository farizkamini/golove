package handler_admin_brand

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handlers struct {
	DBPool *pgxpool.Pool
}

func New(DBPool *pgxpool.Pool) *Handlers {
	return &Handlers{
		DBPool: DBPool,
	}
}
func (h *Handlers) Controller(r chi.Router) {
	r.Route("/trys", func(r chi.Router) {
		r.Get("/", h.Find)
		r.Post("/create", h.Create)
	})
}

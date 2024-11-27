package handler_admin_brand

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handlers struct {
	DB *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Handlers {
	return &Handlers{
		DB: db,
	}
}
func Controller(r chi.Router, db *pgxpool.Pool) {
	h := New(db)
	r.Route("/trys", func(r chi.Router) {
		r.Get("/", h.Find)
		r.Post("/create", h.Create)
	})
}

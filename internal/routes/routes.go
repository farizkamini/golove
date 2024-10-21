package routes

import (
	route_admin "github.com/farizkamini/golove/internal/modules/admin"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Master(DB *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()
	route_admin.Controller(r, DB)
	return r
}

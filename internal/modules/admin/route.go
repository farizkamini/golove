package route_admin

import (
	handler_admin_brand "github.com/farizkamini/golove/internal/modules/admin/handler"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Controller(r chi.Router, DB *pgxpool.Pool) {
	r.Route("/admin", func(r chi.Router) {
		handler_admin_brand.Controller(r, DB)
	})
}

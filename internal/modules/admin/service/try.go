package service_admin_auth

import (
	"net/http"

	repo_admin_auth "github.com/farizkamini/golove/internal/modules/admin/repo"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (s *Service) Try(DBPool *pgxpool.Pool, r *http.Request) ([]repo_admin_auth.Res, error, int) {
	res, err := repo_admin_auth.NewRead(DBPool).Try(s.Ctx)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}
	return res, nil, http.StatusOK
}
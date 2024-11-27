package service_admin_auth

import (
	"net/http"

	repo_admin_auth "github.com/farizkamini/golove/internal/modules/admin/repo"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (s *Service) Create(
	DB *pgxpool.Pool,
	r *http.Request,
) (err error, status int) {
	TX, err := DB.Begin(s.Ctx)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	err = repo_admin_auth.NewWrite(TX).Create(s.Ctx, repo_admin_auth.CreateReq{
		Name: "foo",
		Age:  10,
	})
	if err != nil {
		TX.Rollback(s.Ctx)
		return err, http.StatusInternalServerError
	}
	err = TX.Commit(s.Ctx)
	if err != nil {
		TX.Rollback(s.Ctx)
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusOK
}

package service_admin_auth

import (
	"net/http"

	"github.com/farizkamini/golove/config/db"
	repo_admin_auth "github.com/farizkamini/golove/internal/modules/admin/repo"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (s *Service) Create(
	DBPool *pgxpool.Pool,
	r *http.Request,
) (err error, status int) {
	conn, err := DBPool.Acquire(s.Ctx)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	TX, err := conn.Begin(s.Ctx)
	if err != nil {
	}
	defer conn.Conn().Close(s.Ctx)
	defer conn.Release()
	defer TX.Conn().Close(s.Ctx)
	err = repo_admin_auth.NewWrite(TX).Create(s.Ctx, repo_admin_auth.CreateReq{
		Name: "foo",
		Age:  10,
	})
	err = db.RollCommit(err, s.Ctx, TX)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

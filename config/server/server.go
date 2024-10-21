package server

import (
	"fmt"

	"net/http"

	"strconv"

	"github.com/farizkamini/golove/internal/routes"
	"github.com/farizkamini/golove/pkg/vip"
	"github.com/farizkamini/golove/pkg/zlog"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SrvConfig struct {
	Port    string
	Handler http.Handler
	DBPool  *pgxpool.Pool
}

func New(DBPool *pgxpool.Pool) *SrvConfig {
	return &SrvConfig{
		DBPool: DBPool,
	}
}

func (s *SrvConfig) Serve() {
	vipp, errVip := vip.New().App()
	if errVip != nil {
		zlog.Fatal(errVip)
		return
	}

	srv := &SrvConfig{
		Port:    ":" + strconv.Itoa(vipp.AppPort),
		Handler: routes.Master(s.DBPool),
	}

	zlog.Info(nil, fmt.Sprintf("lets rawk at %s", srv.Port))
	err := http.ListenAndServe(srv.Port, srv.Handler)
	if err != nil {
		zlog.Fatal(err)
	}
}

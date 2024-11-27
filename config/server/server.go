package server

import (
	"context"
	"fmt"

	"net/http"

	"strconv"

	"github.com/farizkamini/golove/config/db"
	"github.com/farizkamini/golove/internal/routes"
	"github.com/farizkamini/golove/pkg/vip"
	"github.com/farizkamini/golove/pkg/zlog"
)

type SrvConfig struct {
	Port    string
	Handler http.Handler
}

func New() *SrvConfig {
	return &SrvConfig{}
}

func (s *SrvConfig) Serve(ctx context.Context) error {
	vipp, errVip := vip.New().App()
	if errVip != nil {
		return errVip
	}
	DB, err := db.New(ctx).Conn()
	if err != nil {
		return err
	}
	defer DB.Close()
	srv := &SrvConfig{
		Port:    ":" + strconv.Itoa(vipp.AppPort),
		Handler: routes.Master(DB),
	}

	zlog.Info(nil, fmt.Sprintf("lets rawk at %s", srv.Port))
	return http.ListenAndServe(srv.Port, srv.Handler)

}

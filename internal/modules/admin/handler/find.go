package handler_admin_brand

import (
	"context"
	"net/http"
	"time"

	service_admin_auth "github.com/farizkamini/golove/internal/modules/admin/service"
	"github.com/farizkamini/golove/pkg/resp"
	"github.com/farizkamini/golove/pkg/zlog"
)

func (h *Handlers) Find(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
	defer cancel()
	res, err, stat := service_admin_auth.New(ctx).Find(h.DB, r)
	if err != nil {
		zlog.Error(err)
		resp.Error(err, stat, w)
		return
	}
	resp.Success(res, "oke", stat, w)
}

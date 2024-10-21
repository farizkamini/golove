package handler_admin_brand

import (
	"net/http"

	service_admin_auth "github.com/farizkamini/golove/internal/modules/admin/service"
	"github.com/farizkamini/golove/pkg/resp"
	"github.com/farizkamini/golove/pkg/zlog"
)

func (h *Handlers) Find(w http.ResponseWriter, r *http.Request) {
	res, err, stat := service_admin_auth.New(r.Context()).Try(h.DBPool, r)
	if err != nil {
		zlog.Error(err)
		resp.Error(err, stat, w)
		return
	}
	resp.Success(res, "oke", stat, w)
}

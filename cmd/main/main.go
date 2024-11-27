package main

import (
	"context"
	"fmt"
	"time"

	"net/http"
	"strconv"

	"github.com/farizkamini/golove/config/server"
	"github.com/farizkamini/golove/config/serverstatic"
	"github.com/farizkamini/golove/pkg/vip"
	"github.com/farizkamini/golove/pkg/zlog"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	vipp, errVip := vip.New().App()
	if errVip != nil {
		zlog.Fatal(errVip)
		return
	}
	go func() {
		zlog.Info(fmt.Sprintf("port asset: %d", vipp.AppPortAsset), "server asset run")
		err := http.ListenAndServe(":"+strconv.Itoa(vipp.AppPortAsset), serverstatic.Master())
		if err != nil {
			zlog.Fatal(err)
			return
		}
	}()

	server.CreateDirAssets()
	err := server.New().Serve(ctx)
	if err != nil {
		zlog.Fatal(err)
		return
	}
}

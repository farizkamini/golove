package main

import (
	"context"
	"fmt"

	"log"

	"net/http"
	"strconv"

	"github.com/farizkamini/golove/config/db"
	"github.com/farizkamini/golove/config/server"
	"github.com/farizkamini/golove/config/serverstatic"
	"github.com/farizkamini/golove/pkg/vip"
	"github.com/farizkamini/golove/pkg/zlog"
)

func main() {
	ctx := context.Background()
	DB, err := db.New(ctx).Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	server.CreateLogFile()
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
	ctx = context.Background()
	_, err = db.New(ctx).Conn()
	if err != nil {
		zlog.Fatal(err)
		return
	}
	server.New(DB).Serve()
}

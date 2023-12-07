package main

import (
	"github.com/labstack/echo/v4"
	"github.com/smf8/http-monitor/common"
	"github.com/smf8/http-monitor/db"
	"github.com/smf8/http-monitor/docker"
	"github.com/smf8/http-monitor/handler"
	"github.com/smf8/http-monitor/store"
)

func main() {
	d := db.Setup("http-monitor.db")
	st := store.NewStore(d)
	dck := docker.NewDocker()

	e := echo.New()
	v1 := e.Group("/api")
	h := handler.NewHandler(st, dck)
	h.RegisterRoutes(v1)

	e.HTTPErrorHandler = common.CustomHTTPErrorHandler
	e.Logger.Fatal(e.Start(":9090"))
}

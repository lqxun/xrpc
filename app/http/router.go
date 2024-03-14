package http

import (
	"github.com/lqxun/xrpc/app/api"
	"github.com/lqxun/xrpc/app/http/middleware"
	"github.com/lqxun/xrpc/core/http_server"
)

func Router() *http_server.XMux {
	mux := http_server.NewXMux()
	mux.Use(middleware.WithLogger)
	mux.HandleFunc("/", api.Hello)

	return mux
}

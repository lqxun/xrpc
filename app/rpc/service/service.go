package service

import (
	"github.com/lqxun/xrpc/app/rpc/middleware"
	"github.com/lqxun/xrpc/app/services"
	"github.com/lqxun/xrpc/core/server"
)

func Service() *server.XRpc {
	
	srv := server.NewXRpc()
	srv.Register("hello", services.Hello, middleware.F1, middleware.F2)

	return srv
}

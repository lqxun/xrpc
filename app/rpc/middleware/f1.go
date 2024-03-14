package middleware

import (
	"fmt"
	"github.com/lqxun/xrpc/core/server"
)

func F1(m *server.Middleware) {
	fmt.Println("middleware 1 start")
	m.Next()
	fmt.Println("middleware 1 end")
}

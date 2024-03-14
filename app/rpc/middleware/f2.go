package middleware

import (
	"fmt"
	"github.com/lqxun/xrpc/core/server"
)

func F2(m *server.Middleware) {
	fmt.Println("middleware 2 start")
	m.Next()
	fmt.Println("middleware 2 end")
}

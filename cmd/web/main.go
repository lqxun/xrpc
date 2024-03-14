package main

import (
	"fmt"
	app "github.com/lqxun/xrpc/app/http"
	"github.com/lqxun/xrpc/config"
	"net/http"
	"time"
)

func main() {
	httpConfig := config.NewHttpConfig()
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", httpConfig.Port),
		Handler:      app.Router(),
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	fmt.Println(fmt.Sprintf("Listening on :%d", httpConfig.Port))
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

package api

import (
	"github.com/lqxun/xrpc/app/services"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("path:%s processing...\n", r.URL.Path)

	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	if name == "" {
		name = "world"
	}

	resp, err := services.Hello(name)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(resp))
}

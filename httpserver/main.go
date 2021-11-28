package main

import (
	"github/shangmingyu0x/x/infra"
	"github/shangmingyu0x/x/internal"
	"net/http"
)

func main() {
	serv := infra.NewServer(8888,
		infra.WaitShutdown(60000), //60s 优雅关闭等待时间
	)

	// set up router
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", internal.HealthHandler)
	mux.HandleFunc("/set/response", internal.SetReponseHandler)
	serv.RegisterHandler(mux)

	serv.Run()
}

package main

import (
	"log"

	"github.com/hindenbug/glog/internal/server"
)

func main() {
	srvr := server.NewHttpServer(":8080")
	log.Fatal(srvr.ListenAndServe())
}

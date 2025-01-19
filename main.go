package main

import (
	"atlas/config"
	"atlas/proxy"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config := config.NewConfig()

	/*
		fmt.Println(config.DenyIPList)
		fmt.Println(config.DenyHTTPHeader)
		fmt.Println(config.Backend)
	*/

	proxy := proxy.NewProxy(config.Backend, config.DenyIPList, config.DenyHTTPHeader, config.DenyHTTPBody)
	http.HandleFunc("/", proxy.Server)

	fmt.Println("Starting WaF atlas service..")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

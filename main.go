package main

import (
	"atlas/config"
	"atlas/inspect"
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

	inspect := inspect.NewInspectHTTPRequest(config.DenyIPList, config.DenyHTTPHeader, config.DenyHTTPBody)

	proxy := proxy.NewProxy(config.Backend, inspect)
	http.HandleFunc("/", proxy.Server)

	fmt.Println("Starting WaF atlas service..")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

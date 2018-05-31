package main

import (
	"/user/danielp/DanniDev/gophercon-k8s-go/pkg/routing"
	"log"
	"net/http"
)

//curl -i http://localhost:8000/home
func main() {
	log.Printf("Service is starting ...")

	r := routing.BaseRouter()
	http.ListenAndServe(":8000", r)
}

package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

func main() {
	restful.Add(New())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

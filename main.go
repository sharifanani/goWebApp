package main

import (
	"github.com/sharifanani/goWebApp/Pages"
	"log"
	"net/http"
)

func main() {
	Pages.InitRoutes()
	log.Fatal(http.ListenAndServe(":9909", nil))
}

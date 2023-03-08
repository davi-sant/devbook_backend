package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Get api devBook")
	getRouter := router.ToGenerate(); 
	log.Fatal(http.ListenAndServe(":5000", getRouter))
}

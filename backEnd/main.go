package main

import (
	"fmt"
	"net/http"

	"github.com/SharanM7/BankingWebApp/router"
)

func main() {
	// router
	r := router.Router()
	fmt.Println("App started at port 3000...")
	http.ListenAndServe("localhost:3000", r)
}

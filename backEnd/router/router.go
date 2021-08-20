package router

import (
	"github.com/SharanM7/BankingWebApp/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", middleware.Home)

	return r
}

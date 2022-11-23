package main

import (
	"fmt"
	"net/http"
	core "samsul96maarif/github.com/go-api-app"
	"samsul96maarif/github.com/go-api-app/config"
	"samsul96maarif/github.com/go-api-app/handler"
	"samsul96maarif/github.com/go-api-app/lib/logger"
	"samsul96maarif/github.com/go-api-app/route"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	logger.Init()
}

func main() {
	db, _ := config.NewDB()
	be := core.NewBe(db)
	handler := handler.NewHandler(&be)

	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()
	api1 := api.PathPrefix("/v1").Subrouter()
	apiRoute := route.ApiRoute{
		R:       api1,
		Handler: &handler,
	}

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "pong")
	})

	router.HandleFunc("/dump", func(w http.ResponseWriter, r *http.Request) {
		err := handler.BE.Usecase.CreateSuperAdmin(r.Context())
		fmt.Fprintln(w, err)
		return
	})

	apiRoute.AddAuthRoute()
	apiRoute.AddItemRoute()
	apiRoute.AddSuperAdminUserRoute()

	headerOk := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Authorization"})
	methodOk := handlers.AllowedMethods([]string{"GET", "DELETE", "POST", "PUT"})
	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.CORS(headerOk, methodOk)(router),
	}
	srv.ListenAndServe()
}

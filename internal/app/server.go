package app

import (
	"context"
	"devtask/internal/app/handlers"
	"devtask/internal/app/middleware"
	"devtask/internal/config"
	"devtask/internal/service/info"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const unsecurePort = ":9000"

// server struct holds the repository interface for handlers
type server struct {
	repo handlers.StorageInfo
}

// RunHTTP initializes and starts the HTTP server
func RunHTTP(_ context.Context, service *info.Service, auth config.AuthInfo) {
	implementation := server{repo: service}

	// Channel to listen for termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// Create and configure the router
	router := createRouter(implementation.repo)
	handler := middleware.BasicAuth(middleware.Logger(router), auth.Username, auth.Password)

	// Create and configure the HTTP server
	unsecureServer := &http.Server{
		Addr:    unsecurePort,
		Handler: handler,
	}

	// Goroutine to start the server
	go func() {
		if err := unsecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err, " unsecure")
		}
	}()

	fmt.Println("Сервер запущен!")

	// Wait for termination signal
	<-quit

	fmt.Println("Завершение работы сервера!")

	// Gracefully shutdown the server
	err := unsecureServer.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err, unsecureServer)
	}

	fmt.Println("Работа сервера завершена!")
}

// createRouter initializes the HTTP router with appropriate routes and handlers
func createRouter(implementation handlers.StorageInfo) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/", handlers.Create(implementation)).Methods("POST")
	router.Handle("/", handlers.List(implementation)).Methods("GET")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}", handlers.QueryParamKey), handlers.GetByID(implementation)).Methods("GET")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}", handlers.QueryParamKey), handlers.Update(implementation)).Methods("PUT")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}", handlers.QueryParamKey), handlers.Delete(implementation)).Methods("DELETE")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/info", handlers.QueryParamKey), handlers.GetISTable(implementation)).Methods("GET")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/info", handlers.QueryParamKey), handlers.AddISInfo(implementation)).Methods("POST")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/info", handlers.QueryParamKey), handlers.UpdateISInfo(implementation)).Methods("PUT")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/imgs", handlers.QueryParamKey), handlers.GetImage(implementation)).Methods("GET")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/imgs", handlers.QueryParamKey), handlers.AddImage(implementation)).Methods("POST")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/docs", handlers.QueryParamKey), handlers.GetDocuments(implementation)).Methods("GET")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/docs", handlers.QueryParamKey), handlers.AddDocument(implementation)).Methods("POST")
	return router
}

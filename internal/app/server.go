package app

import (
	"context"
	"devtask/internal/app/handlers"
	"devtask/internal/app/middleware"
	"devtask/internal/config"
	"devtask/internal/service/pvz"
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

//const securePort = ":9001"

type server struct {
	repo handlers.StoragePVZ
}

func RunHTTP(_ context.Context, service *pvz.Service, auth config.AuthInfo) {
	implementation := server{repo: service}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	router := createRouter(implementation.repo)
	handler := middleware.BasicAuth(middleware.Logger(router), auth.Username, auth.Password)

	unsecureServer := &http.Server{
		Addr:    unsecurePort,
		Handler: handler,
	}

	//secureServer := &http.Server{
	//	Addr:    securePort,
	//	Handler: handler,
	//}

	//go func() {
	//	if err := secureServer.ListenAndServeTLS("server.crt", "server.key"); err != nil && !errors.Is(err, http.ErrServerClosed) {
	//		log.Fatal(err, " secure")
	//	}
	//}()

	go func() {
		if err := unsecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err, " unsecure")
		}
	}()

	<-quit

	fmt.Println("Завершение работы сервера!")

	err := unsecureServer.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err, unsecureServer)
	}

	//err = secureServer.Shutdown(context.Background())
	//if err != nil {
	//	log.Fatal(err, unsecureServer)
	//}

	fmt.Println("Работа сервера завершена!")
}

func createRouter(implementation handlers.StoragePVZ) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/", handlers.Create(implementation)).Methods("POST")
	router.Handle("/", handlers.List(implementation)).Methods("GET")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}", handlers.QueryParamKey), handlers.GetByID(implementation)).Methods("GET")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}", handlers.QueryParamKey), handlers.Update(implementation)).Methods("PUT")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}", handlers.QueryParamKey), handlers.Delete(implementation)).Methods("DELETE")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/info", handlers.QueryParamKey), handlers.GetISTable(implementation)).Methods("GET")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/info", handlers.QueryParamKey), handlers.AddISInfo(implementation)).Methods("POST")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/info", handlers.QueryParamKey), handlers.UpdateISInfo(implementation)).Methods("PUT")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/info/res", handlers.QueryParamKey), handlers.AddRes(implementation)).Methods("POST")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/info/res", handlers.QueryParamKey), handlers.GetRes(implementation)).Methods("GET")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/info/res", handlers.QueryParamKey), handlers.UpdateRes(implementation)).Methods("PUT")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/info/adm", handlers.QueryParamKey), handlers.AddAdm(implementation)).Methods("POST")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/info/adm", handlers.QueryParamKey), handlers.GetAdm(implementation)).Methods("GET")
	router.Handle(fmt.Sprintf("/{%s:[0-9]+}/info/adm", handlers.QueryParamKey), handlers.UpdateAdm(implementation)).Methods("PUT")
	return router
}

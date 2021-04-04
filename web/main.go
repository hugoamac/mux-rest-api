package main

import (
	"log"
	"mux-rest-api/infra"
	"mux-rest-api/web/routes"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {

	log.Println("started application...")

	config, _ := infra.LoadVars()

	port := config.AppPort

	conn := infra.NewConnection(config)
	defer conn.Close()

	router := mux.NewRouter().StrictSlash(true)

	//routes
	routes.ContactRoutes(conn.DB(), router)

	managerMiddleware := negroni.New()
	managerMiddleware.Use(negroni.NewLogger())
	managerMiddleware.UseHandler(router)

	logger := log.New(os.Stderr, "Logger", log.Lshortfile)
	server := &http.Server{
		Addr:         port,
		Handler:      managerMiddleware,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		ErrorLog:     logger,
	}

	log.Printf("server listening at port%s", port)

	log.Fatal(server.ListenAndServe())
}

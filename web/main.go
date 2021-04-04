package main

import (
	"log"
	"mux-rest-api/infra"
	"mux-rest-api/web/routes"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {

	log.Println("started application...")

	infra.LoadVars()

	port := viper.Get("APP_PORT").(string)

	conn := infra.NewConnection()
	defer conn.Close()

	router := mux.NewRouter().StrictSlash(true)

	//routes
	routes.ContactRoutes(conn.DB(), router)

	logger := log.New(os.Stderr, "Logger", log.Lshortfile)
	server := &http.Server{
		Addr:         port,
		Handler:      router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		ErrorLog:     logger,
	}

	log.Printf("server listening at port%s", port)

	log.Fatal(server.ListenAndServe())
}

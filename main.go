package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/olegsobchuk/moneyrate/config/router"
)

// ServerPort port for server
var ServerPort = os.Getenv("PORT")

func run() error {
	router := router.MakeRouter()
	server := &http.Server{
		Handler:      router,
		Addr:         ":" + ServerPort,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Start server...")
	log.Println("Start server. Port ", ServerPort)
	log.Fatal(run())
}

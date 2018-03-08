package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/olegsobchuk/moneyrate/config/router"
)

func run() error {
	router := router.MakeRouter()
	server := &http.Server{
		Handler:      router,
		Addr:         ":3001",
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
	log.Println("Start server...")
	log.Fatal(run())
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var app AppConfig = *NewAppConfig()

func main() {
	app.DoneChan = make(chan bool)
	parseFlags(&app)

	srv := &http.Server{
		Addr:              app.Port,
		Handler:           routes(),
		IdleTimeout:       time.Second * 30,
		ReadTimeout:       time.Second * 10,
		ReadHeaderTimeout: time.Second * 25,
		WriteTimeout:      time.Second * 10,
	}

	fmt.Printf("Starting application on %s\n", srv.Addr)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go app.UpdateTimeStamp()
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

		<-sigChan
		log.Println("Received shutdown signal...")

		ctx.Done()
		cancel()
		app.DoneChan <- true

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP server shutdown error: %v", err)
		}
	}()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v", err)
			app.DoneChan <- true
		}
	}()
	<-app.DoneChan
}

func parseFlags(a *AppConfig) {

	flag.StringVar(&a.Port, "port", ":8080", "the port the application runs on")
	flag.StringVar(&a.BaseURL, "base_url", "/", "the base url of the application")

	flag.Parse()

}

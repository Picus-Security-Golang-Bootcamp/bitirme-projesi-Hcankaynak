package server

import (
	"context"
	"fmt"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServer(cfg *config.ServerConfig) {
	// TODO: relase mode?
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	srv := &http.Server{
		Addr:         fmt.Sprintf("localhost:%s", cfg.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSecs * int64(time.Second)),
		WriteTimeout: time.Duration(cfg.WriteTimeoutSecs * int64(time.Second)),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		log.Println("Book store service started")
	}()

	shutdownGin(srv, time.Duration(cfg.TimeoutSecs*int64(time.Second)))
}

func shutdownGin(srv *http.Server, timeout time.Duration) {

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of seconds given in config.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down BOSS server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("BOSS server forced to shutdown: ", err)
	}

	log.Println("BOSS server exiting")
}

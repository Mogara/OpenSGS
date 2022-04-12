package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Mogara/OpenSGS/pkg/apiserver"
)

var (
	host     string
	port     int
	logLevel string
)

func init() {
	flag.StringVar(&host, "host", "localhost", "host to listen on")
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.StringVar(&logLevel, "log-level", "info", "log level")
	flag.Parse()

	log.SetOutput(os.Stdout)
	lvl, err := log.ParseLevel(logLevel)
	if err != nil {
		log.WithError(err).Fatalf("invalid log level")
	}
	log.SetLevel(lvl)
}

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	s := apiserver.NewAPIServer(host, port)
	if err := s.PrepareRun(); err != nil {
		panic(err)
	}
	go func() {
		if err := s.Run(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Infof("Start listening on %s", s.Server.Addr)

	<-done

	log.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()
	if err := s.Server.Shutdown(ctx); err != nil {
		log.WithError(err).Warnf("Server shutdown failed")
	} else {
		log.Info("Server gracefully stopped")
	}
}

package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Mogara/OpenSGS/pkg/apiserver"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host", "localhost", "host to listen on")
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.Parse()
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
	fmt.Println("Start listening on", s.Server.Addr)

	<-done

	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()
	if err := s.Server.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown failed: %v\n", err)
	} else {
		fmt.Println("Server gracefully stopped")
	}
}

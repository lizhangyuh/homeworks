package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	exit := make(chan struct{})

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	mux.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("server get exit cmd")
		exit <- struct{}{}
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// shutdown
	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Printf("errgroup canceled\n")
		case <-exit:
			log.Printf("exit by request\n")
		}

		// 等待5s后退出
		log.Printf("server will exit in 5 sec...\n")
		timeoutCxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		return server.Shutdown(timeoutCxt)
	})

	// server
	g.Go(func() error {
		log.Printf("server start")
		return server.ListenAndServe()
	})

	// os exit
	g.Go(func() error {
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			log.Printf("os exit cancel")
			return nil
		case sign := <-s:
			fmt.Printf("%s\n", sign)
			return fmt.Errorf("os exit")
		}
	})

	log.Printf("all exit by: %s", g.Wait())

}

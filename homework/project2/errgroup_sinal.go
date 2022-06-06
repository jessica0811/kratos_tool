/*
 * @Author: qiuhua.lin
 * @Date: 2022-06-05 16:33:59
 */
package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	serverOut := make(chan struct{})
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, GopherCon SG")
	})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- struct{}{}
	})
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	g.Go(func() error {
		return server.ListenAndServe()
	})
	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup exit....")
		case <-serverOut:
			log.Println("server will out....")
		}
		timeOutCtx, _ := context.WithTimeout(context.Background(), 3*time.Second)
		return server.Shutdown(timeOutCtx)
	})
	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("get os signal: %v", sig)
		}
	})
	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}

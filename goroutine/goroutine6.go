// /*
//  * @Author: qiuhua.lin
//  * @Date: 2022-06-16 18:37:22
//  */
// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/pkg/errors"
// 	"golang.org/x/sync/errgroup"
// )

// func main() {
// 	g, ctx := errgroup.WithContext(context.Background())
// 	serverOut := make(chan struct{})
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Hello, GopherCon SG")
// 	})
// 	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("----进来这里---")
// 		serverOut <- struct{}{}  // 为空会阻塞
// 	})
// 	server := http.Server{
// 		Handler: mux,
// 		Addr:    ":8080",
// 	}
// 	g.Go(func() error {
// 		return server.ListenAndServe()
// 	})
// 	g.Go(func() error {
// 		select {
// 		case <-ctx.Done():
// 			log.Println("errgroup exit....")
// 		case <-serverOut:
// 			log.Println("server will out....")
// 		}
// 		log.Println("这里了")
// 		timeOutCtx, _ := context.WithTimeout(context.Background(), 3*time.Second)
// 		return server.Shutdown(timeOutCtx)
// 	})
// 	g.Go(func() error {
// 		quit := make(chan os.Signal, 0)
// 		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 		select {
// 		case <-ctx.Done():
// 			return ctx.Err()
// 		case sig := <-quit:
// 			fmt.Println("先来这里")
// 			return errors.Errorf("get os signal: %v", sig)
// 		}
// 	})
// 	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
// }

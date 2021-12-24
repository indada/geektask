package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	g,ctx := errgroup.WithContext(context.Background())
	stop := make(chan struct{})
	g.Go(func() error{
		return serverApp(ctx,stop)
	})
	g.Go(func() error{
		return serverDebug(ctx,stop)
	})
	g.Go(func() error { //监听信号
		sign := make(chan os.Signal,1)
		//监听信号 kill ctrl+c 信号值传递给通道sign
		signal.Notify(sign,syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,syscall.SIGQUIT)
		select {
		case <-sign:
			//取出信号
			fmt.Println("Notify")
			return errors.New("notify")
		case <-stop:
			fmt.Println("Signal stop bay")
		}
		defer close(sign)
		return nil
	})
	err := g.Wait()
	fmt.Println("emo!!",err)
}
func serverDebug(ctx context.Context,stop chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello",hellos)
	mux.HandleFunc("/stop", func(writer http.ResponseWriter, request *http.Request) {
		close(stop)
	})
	return server(":8888",mux,ctx,stop)
}
func serverApp(ctx context.Context,stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello",hello)
	return server(":7777",mux,ctx,stop)
}
func server(addr string,mux http.Handler,ctx context.Context,stop <-chan struct{}) error { //g errgroup.Group
	ser := http.Server{
		Addr: addr,
		Handler: mux,
	}
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("退出:",addr)
			ser.Shutdown(context.Background())
		case <-stop:
			fmt.Println("stop退出:",addr)
			ser.Shutdown(context.Background())
		}
	}()
	return ser.ListenAndServe()
}
func hello(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintln(w,"hello7777,GopherCon")
}
func hellos(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintln(w,"hellos8888,GopherCon")
}
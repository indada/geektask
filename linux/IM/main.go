package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"fmt"
	"net/http"
	"html/template"
	"os"
	"os/signal"
	"syscall"
	"github.com/gorilla/websocket"
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
		case <-ctx.Done():
			fmt.Println("ctx err")
			return errors.New("ctx err")
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
	mux.HandleFunc("/",echo)
	mux.HandleFunc("/stop", func(writer http.ResponseWriter, request *http.Request) {
		close(stop)
	})
	return server(":88",mux,ctx,stop)
}
func serverApp(ctx context.Context,stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello",hello)
	mux.HandleFunc("/",im)
	return server(":80",mux,ctx,stop)
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
func im(w http.ResponseWriter,r *http.Request)  {
	// 解析模板
	t,err := template.ParseFiles("./html/index.html")
	if err!=nil {
		fmt.Println("Parse template failed, err%v", err)
		return
	}
	data := "hello"
	err = t.Execute(w,data)
	if err!=nil {
		fmt.Println("execute template failed, err%v", err)
		return
	}
	//fmt.Fprintln(w,"即时通信！")
}
func echo(w http.ResponseWriter,r *http.Request)  {
	//先升级http协议，此处nil可以填写其他内容，本次只是为了简单应用websocket
	//nil可以填cookie和其他实现内容
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}
	defer c.Close()
	count := 0
	for {
		if count += 1; count > 100{
			break
		}
		//建立信息
		message := "1111"
		//log.Printf("recv: %s", message)
		//此处1指代后面的内容为字符串  发送信息给客户端
		err = c.WriteMessage(1, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
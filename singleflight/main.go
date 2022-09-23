package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"golang.org/x/sync/singleflight"
)

var count int32
func main() {
	//只执行一次
	time.AfterFunc(1*time.Second, func() {
		atomic.AddInt32(&count, -count)
	})
	var wg sync.WaitGroup
	new := time.Now()
	n := 1000
	var sg = &singleflight.Group{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			res,_ := singleflightGetArticle(sg,i)
			if res != "ss" {

			}
			//_,_ = getArticle(i)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("同时发起 %d 次请求，count: %d，耗时 %s ",n,count,time.Since(new))
}
func singleflightGetArticle(sg *singleflight.Group,id int) (string,error) {
	v,err,_ := sg.Do(fmt.Sprintf("%d", id), func() (interface{}, error) {
		return getArticle(id)
	})
	return v.(string),err
}
func getArticle(id int) (article string,err error) {
	atomic.AddInt32(&count,1)
	time.Sleep(time.Duration(count)*time.Millisecond)
	fmt.Println("1")
	return fmt.Sprintf("article %d",id),nil
}

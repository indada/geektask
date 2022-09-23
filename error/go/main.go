package  _go

//封装goroutine 预防panic导致主进程退出
func Go(f func())  {
	go func() {
		defer func() {
			if err:= recover();err!=nil {
				log.Printf("panic:%+v",err)
			}
		}()

		f()
	}()
}

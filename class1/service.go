package main

import (
	"net/http"
)

func serve(addr string, handler http.Handler) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return s.ListenAndServe()
}

//func serveApp() error {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
//		fmt.Fprintf(writer, "hello, QCon!")
//	})
//	if err := serve("0.0.0.0:8080", mux); err != nil {
//		return errors.New("服务出问题啦")
//
//	}
//	return nil
//}

//func main()  {
//	done := make(chan error, 1)
//	stop := make(chan struct{})
//
//	go func() {
//		done <- serveApp()
//	}()
//	var stopped bool
//	for i := 0; i < cap(done); i ++ {
//		if err := <-done; err != nil {
//            fmt.Println("error: %v", err)
//		}
//		if !stopped {
//			stopped = true
//			close(stop)
//		}
//	}
//}

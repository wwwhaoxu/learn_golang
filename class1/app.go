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

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

type App struct {
	addr    string
	ctx     context.Context
	cancel  func()
	handler http.Handler
}

func New(addr string, handler http.Handler) *App {

	ctx, cancle := context.WithCancel(context.Background())

	return &App{
		addr:    addr,
		ctx:     ctx,
		cancel:  cancle,
		handler: handler,
	}
}

func (a *App) Run() error {

	g, ctx := errgroup.WithContext(a.ctx)
	for {
		g.Go(func() error {
			return serve(a.addr, a.handler)
		})
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, shutdownSignals...)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				a.stop()
			}
		}
	})
	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (a *App) stop() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello, QCon!")
	})
	app := New("0.0.0.0:8080", mux)
	time.AfterFunc(time.Second, func() {
		app.stop()
	})
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}

}

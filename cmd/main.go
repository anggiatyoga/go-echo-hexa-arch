package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/anggiatyoga/hris-api/cmd/cli"
	"github.com/anggiatyoga/hris-api/internal/platform/webapi/router"
	"github.com/go-pg/pg"
)

func main() {
	if err := run(); nil != err {
		panic(err.Error())
	}
}

func run() error {
	// GET CONFIG ENV
	conf, err := cli.GetConfig()
	if err != nil {
		fmt.Printf("run app error: %s\n", err.Error())
		return err
	}

	// CONNECT DB
	db := pg.Connect(&pg.Options{
		Addr:     conf.DB.Host,
		User:     conf.DB.User,
		Password: conf.DB.Password,
		Database: conf.DB.Name,
	})
	defer db.Close()
	db.AddQueryHook(Hook{})

	// APP SERVER
	s := &http.Server{
		Addr:         conf.AppConfig.Address,
		ReadTimeout:  conf.AppConfig.ReadTimeout,
		WriteTimeout: conf.AppConfig.WriteTimeout,
	}

	modules, err := cli.Bootstrap(db, conf)
	if err != nil {
		return errors.New(fmt.Sprintf("bootstrapping dependencies: %s", err.Error()))
	}

	e, err := router.Run(modules)
	if err != nil {
		fmt.Printf("running service route error: %s\n", err.Error())
		return err
	}

	go func() {
		if err := e.StartServer(s); err != nil {
			fmt.Printf("shutting down server... : %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), conf.AppConfig.ShutdownTimeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

type Hook struct {
}

func (h Hook) BeforeQuery(e *pg.QueryEvent) {
	s, _ := e.FormattedQuery()
	fmt.Printf("query: %s", s)
}

func (h Hook) AfterQuery(e *pg.QueryEvent) {

}

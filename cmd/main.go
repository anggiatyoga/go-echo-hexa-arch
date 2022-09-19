package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/anggiatyoga/hris-api/cmd/cli"
	"github.com/anggiatyoga/hris-api/internal/platform/webapi/router"
	_ "github.com/lib/pq"
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
	// db := pg.Connect(&pg.Options{
	// 	Addr:     fmt.Sprintf("%s:%s", conf.DB.Host, conf.DB.Port),
	// 	User:     conf.DB.User,
	// 	Password: conf.DB.Password,
	// 	Database: conf.DB.Name,
	// })
	// defer db.Close()
	// db.AddQueryHook(Hook{})

	var db *sql.DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s sslmode=disable",
		conf.DB.Host, conf.DB.User, conf.DB.Password, conf.DB.Name, conf.DB.Port, conf.DB.Schema)

	db, err = sql.Open(conf.DB.Driver, dsn)
	if err != nil {
		info := fmt.Sprintf("Unable to connect to database: %v\n", err.Error())
		fmt.Printf(info)
		panic("connectionString error")
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("DSN Invalid: ", err.Error())
		panic("DSN Invalid")
	}

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

// type Hook struct {
// }

// func (h Hook) BeforeQuery(e *pg.QueryEvent) {
// 	s, _ := e.FormattedQuery()
// 	fmt.Printf("query: %s\n", s)
// }

// func (h Hook) AfterQuery(e *pg.QueryEvent) {

// }

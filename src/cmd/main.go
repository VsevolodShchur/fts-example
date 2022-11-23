package main

import (
	"context"
	"log"
	"menu/internal"
	"menu/pkg/manticore"
	"menu/pkg/server"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql" // Manticore Search
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres
)

func main() {
	db, err := sqlx.Connect("postgres", "host=db port=5432 user=postgres password=password dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	manticoreClient, err := manticore.NewClient("any@tcp(manticore:9306)/any")
	if err != nil {
		log.Fatal(err)
	}
	defer manticoreClient.Close()

	app := internal.InitNewApp(db, manticoreClient)
	appRouter := app.GetRouter()

	srv := server.NewServer("8000", appRouter)

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}

package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/lolitsgab/dist-kv-store/application"
)

func main() {
	app := application.New()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}

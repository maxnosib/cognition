package main

import (
	"context"
	"log"

	"github.com/maxnosib/cognition/internal/config"
	"github.com/maxnosib/cognition/internal/handlers"
	"github.com/maxnosib/cognition/internal/repository"
	"github.com/maxnosib/cognition/internal/usecase"
)

func main() {
	ctx := context.Background()
	cfgServer := config.NewServer()
	if err := cfgServer.Validate(); err != nil {
		log.Fatal(err)
	}

	cfgDB, err := repository.New()
	if err != nil {
		log.Fatal(err)
	}

	uc := usecase.New(nil, cfgDB, cfgDB, cfgDB)

	uc.Categories, err = uc.GetCategories(ctx)
	if err != nil {
		log.Fatal(err)
	}

	handlers.New(ctx, cfgServer, uc)
}

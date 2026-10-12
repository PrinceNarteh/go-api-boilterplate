package main

import (
	"log"

	"github.com/PrinceNarteh/go-api/internal/db"
	"github.com/PrinceNarteh/go-api/internal/repositories"
	"github.com/PrinceNarteh/go-api/internal/services"
)

func main() {
	// connecting to database
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	defer func() {
		if err := db.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}()

	repo := repositories.NewRepositories(db.Database)
	_ = services.NewServices(repo)

	_ = initApp()
}

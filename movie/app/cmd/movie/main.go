package main

import (
	"flag"
	"fmt"
	"log"

	movieRepo "app/internal/repository"

	"app/internal/delivery/http"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found!")
	}
}

func seedMovies(s *movieRepo.PostgresStore) {
	_, err := s.CreateMovie("Batman and Robin", 1949, "batman-and-robin")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	store, err := movieRepo.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(nil); err != nil {
		log.Fatal(err)
	}

	seed := flag.Bool("seed", false, "Seed the DB")

	if *seed {
		fmt.Println("seeding the database")
		seedMovies(store)
	}

	server := http.NewAPIServer(":3000", store)
	server.Run()
}

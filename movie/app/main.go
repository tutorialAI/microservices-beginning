package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func seedMovies(s *PostgresStore) {
	err := s.CreateMovie("Batman and Robin", 1949, "batman-and-robin")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	seed := flag.Bool("seed", false, "Seed the DB")

	if *seed {
		fmt.Println("seeding the database")
		seedMovies(store)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}

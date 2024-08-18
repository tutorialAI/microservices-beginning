package main

import (
	"database/sql"
	"fmt"
	"os"

	"errors"
	"time"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	user, ok := os.LookupEnv("POSTGRES_USER")
	dbName, ok := os.LookupEnv("POSTGRES_DB")
	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	port, ok := os.LookupEnv("POSTGRES_PORT")

	if !ok {
		return nil, errors.New("cannot get env configs for db connection")
	}

	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s port=%s sslmode=disable",
		user, dbName, password, port,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createMoviesTable()
}

func (s *PostgresStore) createMoviesTable() error {
	query := `create table if not exists movies (
		id serial primary key,
		title varchar(250),
		link varchar(250),
		release_year timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) GetMovies() ([]*Movie, error) {
	rows, err := s.db.Query("select * from movies")

	if err != nil {
		return nil, err
	}

	movies := []*Movie{}

	if rows.Next() {
		movie, err := scanIntoMovie(rows)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (s *PostgresStore) GetMovieByLink(link string) (*Movie, error) {
	rows, err := s.db.Query("select * from movies where link = $1", link)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoMovie(rows)
	}

	return nil, fmt.Errorf("movie %s not found", link)
}

func (s *PostgresStore) CreateMovie(title string, releaseYear int, link string) error {
	date := time.Date(releaseYear, 11, 17, 20, 34, 58, 651387237, time.UTC)

	_, err := s.db.Exec(
		"insert into movies(title, release_year, link) values($1, $2, $3)",
		title, date, link,
	)

	if err != nil {
		return err
	}

	return nil
}

func scanIntoMovie(rows *sql.Rows) (*Movie, error) {
	movie := new(Movie)
	err := rows.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Link,
		&movie.ReleaseYear,
	)

	return movie, err
}

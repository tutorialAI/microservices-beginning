package repository

import (
	"database/sql"
	"fmt"
	"os"

	"app/internal/models"
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

func (s *PostgresStore) Init(db *sql.DB) error {
	if db != nil {
		s.db = db
	}
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

func (s *PostgresStore) GetMovies() ([]*models.Movie, error) {
	rows, err := s.db.Query("select * from movies")

	if err != nil {
		return nil, err
	}

	movies := []*models.Movie{}

	if rows.Next() {
		movie, err := scanIntoMovie(rows)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (s *PostgresStore) GetMovieByLink(link string) (*models.Movie, error) {
	rows, err := s.db.Query("select * from movies where link = $1", link)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoMovie(rows)
	}

	return nil, fmt.Errorf("movie %s not found", link)
}

func (s *PostgresStore) CreateMovie(title string, releaseYear int, link string) (int, error) {
	date := time.Date(releaseYear, 11, 17, 20, 34, 58, 651387237, time.UTC)

	var movieID int
	err := s.db.QueryRow(
		"insert into movies(title, release_year, link) values($1, $2, $3) returning id",
		title, date, link,
	).Scan(&movieID)

	if err != nil {
		return 0, err
	}

	return movieID, nil
}

func (s *PostgresStore) AddLike(movieID int, userID int) (done bool, err error) {
	tx, err := s.db.Begin()

	if err != nil {
		return false, errors.New("cannot start transaction")
	}

	defer func() {
		if err != nil {
			errRb := tx.Rollback()
			if errRb != nil {
				err = errors.New("error during rollback")
				return
			}

			return
		}

		err = tx.Commit()
	}()

	_, err = s.db.Exec("set transaction isolation level repatable read;")
	if err != nil {
		return false, errors.New("set transaction level")
	}

	return s.addLike(movieID, userID)
}

func (s *PostgresStore) UpdateMovie(movieID int, title string, releaseYear int, link string) (bool, error) {
	_, err := s.db.Exec("update movies set title = $1, release_year = $2, link = $3 where id = $4",
		title, releaseYear, link, movieID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *PostgresStore) CloseDB() {
	s.db.Close()
}

func (s *PostgresStore) addLike(movieID int, userID int) (bool, error) {
	res, err := s.db.Exec(
		"insert into user_likes (user_id, movie_id) values($1, $2) on conflict do nothing;",
		movieID, userID,
	)

	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	res, err = s.db.Exec("update movies set likes = likes+1 where id = $1;", movieID)

	if err != nil {
		return false, err
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, errors.New("likes nott updated")
	}

	return true, nil
}

func scanIntoMovie(rows *sql.Rows) (*models.Movie, error) {
	movie := new(models.Movie)
	err := rows.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Link,
		&movie.ReleaseYear,
	)

	return movie, err
}

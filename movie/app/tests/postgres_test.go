package main

import (
	"database/sql"
	"testing"
	"time"

	movieRepo "app/internal/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetMovies(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	store := &movieRepo.PostgresStore{}
	store.Init(db)

	t.Run("empty database", func(t *testing.T) {
		mock.ExpectQuery("select \\* from movies").WillReturnRows(sqlmock.NewRows(nil))
		movies, err := store.GetMovies()
		assert.NoError(t, err)
		assert.Empty(t, movies)
	})

	t.Run("single movie in database", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "link", "release_year"}).
			AddRow(1, "Movie 1", "link1", time.Now())
		mock.ExpectQuery("select \\* from movies").WillReturnRows(rows)
		movies, err := store.GetMovies()
		assert.NoError(t, err)
		assert.Len(t, movies, 1)
		assert.Equal(t, "Movie 1", movies[0].Title)
	})

	t.Run("multiple movies in database", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "link", "release_year"}).
			AddRow(1, "Movie 1", "link1", time.Now()).
			AddRow(2, "Movie 2", "link2", time.Now())
		mock.ExpectQuery("select \\* from movies").WillReturnRows(rows)
		movies, err := store.GetMovies()
		assert.NoError(t, err)
		assert.Len(t, movies, 2)
		assert.Equal(t, "Movie 1", movies[0].Title)
		assert.Equal(t, "Movie 2", movies[1].Title)
	})

	t.Run("database query error", func(t *testing.T) {
		mock.ExpectQuery("select \\* from movies").WillReturnError(sql.ErrConnDone)
		movies, err := store.GetMovies()
		assert.Error(t, err)
		assert.Empty(t, movies)
	})
}

func TestUpdateMovie(t *testing.T) {
	db, err := sql.Open("postgres", "user=myuser dbname=mydb sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	store := &movieRepo.PostgresStore{}
	store.Init(db)

	// Successful update of a movie
	movieID := 1
	title := "New Title"
	releaseYear := 2022
	link := "https://example.com"
	updated, err := store.UpdateMovie(movieID, title, releaseYear, link)
	assert.True(t, updated)
	assert.NoError(t, err)

	// Update of a non-existent movie
	nonExistentMovieID := 999
	updated, err = store.UpdateMovie(nonExistentMovieID, title, releaseYear, link)
	assert.False(t, updated)
	assert.NoError(t, err)

	// Update with invalid input (e.g. empty title)
	emptyTitle := ""
	updated, err = store.UpdateMovie(movieID, emptyTitle, releaseYear, link)
	assert.False(t, updated)
	assert.Error(t, err)

	// Update with database error
	store.CloseDB()
	updated, err = store.UpdateMovie(movieID, title, releaseYear, link)
	assert.False(t, updated)
	assert.Error(t, err)
}

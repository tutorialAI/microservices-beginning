package repository

import "app/internal/models"

type UserRepository interface {
	Init() error
	GetMovies() ([]*models.Movie, error)
	GetMovieByLink(link string) (*models.Movie, error)
	CreateMovie(title string, releaseYear int, link string) (int, error)
	AddLike(movieID int, userID int) (bool, error)
	UpdateMovie(title string, releaseYear int, link string, movieID int) error
}

package main

type Storage interface {
	// CreateAccount(*Account) error
	// DeleteAccount(int) error
	// UpdateAccount(*Account) error

	CreateMovie(title string, releaseYear int, link string) error
	GetMovies() ([]*Movie, error)
	GetMovieByLink(string) (*Movie, error)
}

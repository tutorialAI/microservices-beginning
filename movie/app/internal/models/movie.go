package models

type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	ReleaseYear int    `json:"releaseYear"`
	Link        string `json:"link"`
}

type CreateMovieRequest struct {
	Title       string `json:"title"`
	ReleaseYear int    `json:"releaseYear"`
	Link        string `json:"link"`
}

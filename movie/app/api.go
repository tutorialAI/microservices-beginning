package main

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	// router.HandleFunc("/login", makeHTTPHandleFunc(s.handleLogin))
	router.HandleFunc("/movies", makeHTTPHandleFunc(s.handleMovies))
	router.HandleFunc("/movies/{link}", makeHTTPHandleFunc(s.GetMovieByLink))

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleMovies(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		return s.GetMovies(w, r)
	}

	if r.Method == "POST" {
		return s.CreateMovie(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) CreateMovie(w http.ResponseWriter, r *http.Request) error {
	req := new(CreateMovieRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	if err := s.store.CreateMovie(req.Title, req.ReleaseYear, req.Link); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, req.Title+"was created succesfully")
}

func (s *APIServer) GetMovies(w http.ResponseWriter, r *http.Request) error {
	movies, err := s.store.GetMovies()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, movies)
}

func (s *APIServer) GetMovieByLink(w http.ResponseWriter, r *http.Request) error {
	link := mux.Vars(r)["link"]
	movie, err := s.store.GetMovieByLink(link)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, movie)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

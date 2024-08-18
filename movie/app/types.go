package main

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

// func NewAccount(firstName, lastName, password string) (*Account, error) {
// 	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Account{
// 		FirstName:         firstName,
// 		LastName:          lastName,
// 		EncryptedPassword: string(encpw),
// 		Number:            int64(rand.Intn(1000000)),
// 		ReleaseYear:         time.Now().UTC(),
// 	}, nil
// }

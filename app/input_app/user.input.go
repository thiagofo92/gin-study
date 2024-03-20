package inputapp

type UserInput struct {
	Name         string   `json:"name"`
	Password     string   `json:"password"`
	Email        string   `json:"email"`
	RentedBooks  []string `json:"rentedBooks"`
	BooksHistory []string `json:"booksHistory"`
}

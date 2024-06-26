package schema

type UserSchema struct {
	Id           string   `bson:"_id,omitempty"`
	Name         string   `bson:"name,omitempty"`
	Password     string   `bson:"password,omitempty"`
	Email        string   `bson:"email,omitempty"`
	RentedBooks  []string `bson:"rentedBooks,omitempty"`
	BooksHistory []string `bson:"booksHistory,omitempty"`
}

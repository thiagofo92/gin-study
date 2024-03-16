package schema

type UserSchema struct {
	Id           string   `bson:"_id,omitempty"`
	Name         string   `bson:"name"`
	Password     string   `bson:"password"`
	Email        string   `bson:"email"`
	RentedBooks  []string `bson:"rentedBooks"`
	BooksHistory []string `bson:"booksHistory"`
}

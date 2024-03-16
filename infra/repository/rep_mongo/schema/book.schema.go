package schema

type BookSchema struct {
	Id         string   `bson:"_id,omitempty"`
	Name       string   `bson:"name"`
	Author     string   `bson:"author"`
	Categories []string `bson:"categories"`
	Available  uint16   `bson:"available"`
	Rented     uint16   `bson:"rented"`
}

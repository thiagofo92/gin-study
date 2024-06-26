package schema

type BookSchema struct {
	Id         string   `bson:"_id,omitempty"`
	Name       string   `bson:"name,omitempty"`
	Author     string   `bson:"author,omitempty"`
	Categories []string `bson:"categories,omitempty"`
	Available  uint16   `bson:"available,omitempty"`
	Rented     uint16   `bson:"rented,omitempty"`
}

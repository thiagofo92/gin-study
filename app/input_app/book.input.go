package inputapp

type BookInput struct {
	Id         string   `json:"_id,omitempty"`
	Name       string   `json:"name"`
	Author     string   `json:"author"`
	Categories []string `json:"categories"`
	Available  uint16   `json:"available"`
	Rented     uint16   `json:"rented"`
}

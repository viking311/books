package storage

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int16  `json:"year,omitempty"`
	Id     int64  `json:"id,omitempty"`
}

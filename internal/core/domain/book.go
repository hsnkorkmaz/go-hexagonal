package domain

type Book struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Author Author `json:"author"`
}

func NewBook(id int64, name string, author Author) Book {
	return Book{
		Id:     id,
		Name:   name,
		Author: author,
	}
}

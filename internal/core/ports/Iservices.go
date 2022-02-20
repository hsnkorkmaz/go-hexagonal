package ports

import "github.com/hsnkorkmaz/go-hexagonal/internal/core/domain"

type IBookService interface {
	// Book
	CreateBook(book domain.Book) error
	GetBook(id int64) (domain.Book, error)
	GetBooks() ([]domain.Book, error)
	UpdateBook(book domain.Book) error
	DeleteBook(id int64) error
}

type IAuthorService interface {
	// Author
	CreateAuthor(author domain.Author) error
	GetAuthor(id int64) (domain.Author, error)
	GetAuthors() ([]domain.Author, error)
	UpdateAuthor(author domain.Author) error
	DeleteAuthor(id int64) error
}

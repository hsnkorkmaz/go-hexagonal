package services

import (
	"github.com/hsnkorkmaz/go-hexagonal/internal/core/domain"
	"github.com/hsnkorkmaz/go-hexagonal/internal/core/ports"
)

type bookService struct {
	bookRepository ports.IBookRepository
}

func NewBookService(bookRepository ports.IBookRepository) *bookService {
	return &bookService{bookRepository}
}

func (s *bookService) CreateBook(book domain.Book) error {
	return s.bookRepository.CreateBook(book)
}

func (s *bookService) GetBook(id int64) (domain.Book, error) {
	return s.bookRepository.GetBook(id)
}

func (s *bookService) GetBooks() ([]domain.Book, error) {
	return s.bookRepository.GetBooks()
}

func (s *bookService) UpdateBook(book domain.Book) error {
	return s.bookRepository.UpdateBook(book)
}

func (s *bookService) DeleteBook(id int64) error {
	return s.bookRepository.DeleteBook(id)
}
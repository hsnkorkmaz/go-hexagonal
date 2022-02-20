package services

import (
	"github.com/hsnkorkmaz/go-hexagonal/internal/core/domain"
	"github.com/hsnkorkmaz/go-hexagonal/internal/core/ports"
)

type authorService struct {
	authorRepository ports.IAuthorRepository
}

func NewAuthorService(authorRepository ports.IAuthorRepository) *authorService {
	return &authorService{authorRepository}
}

func (s *authorService) CreateAuthor(author domain.Author) error {
	return s.authorRepository.CreateAuthor(author)
}

func (s *authorService) GetAuthor(id int64) (domain.Author, error) {
	return s.authorRepository.GetAuthor(id)
}

func (s *authorService) GetAuthors() ([]domain.Author, error) {
	return s.authorRepository.GetAuthors()
}

func (s *authorService) UpdateAuthor(author domain.Author) error {
	return s.authorRepository.UpdateAuthor(author)
}

func (s *authorService) DeleteAuthor(id int64) error {
	return s.authorRepository.DeleteAuthor(id)
}

package http

import "github.com/hsnkorkmaz/go-hexagonal/internal/core/ports"

type BookHandler struct {
	bookService ports.IBookService
}

func NewBookHandler(bookService ports.IBookService) *BookHandler {
	return &BookHandler{bookService}
}

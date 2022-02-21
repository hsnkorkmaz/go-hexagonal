package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hsnkorkmaz/go-hexagonal/internal/core/domain"
	"github.com/hsnkorkmaz/go-hexagonal/internal/core/ports"
)

type AuthorHandler struct {
	authorService ports.IAuthorService
}

func NewAuthorHandler(authorService ports.IAuthorService) *AuthorHandler {
	return &AuthorHandler{authorService}
}

func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author domain.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.authorService.CreateAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *AuthorHandler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	author, err := a.authorService.GetAuthor(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(author)
}

func (a *AuthorHandler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := a.authorService.GetAuthors()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authors)
}

func (a *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var author domain.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = a.authorService.UpdateAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = a.authorService.DeleteAuthor(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

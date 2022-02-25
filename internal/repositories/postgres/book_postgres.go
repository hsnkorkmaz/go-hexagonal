package postgres

import (
	"database/sql"

	"github.com/hsnkorkmaz/go-hexagonal/internal/core/domain"
	_ "github.com/lib/pq"
)

type bookPostgres struct {
	db *sql.DB
}

func NewBookPostgres(db *sql.DB) *bookPostgres {
	return &bookPostgres{db}
}

func (sql *bookPostgres) CreateBook(book domain.Book) error {
	query := `INSERT INTO books (name, author_id) VALUES ($1, $2)`
	_, err := sql.db.Exec(query, book.Name, book.Author.Id)
	if err != nil {
		return err
	}
	return nil
}

func (sql *bookPostgres) GetBook(id int64) (domain.Book, error) {
	query := `SELECT id, name, author_id FROM books WHERE id = $1`
	rows, err := sql.db.Query(query, id)
	if err != nil {
		return domain.Book{}, err
	}
	defer rows.Close()

	var book domain.Book
	for rows.Next() {
		err := rows.Scan(&book.Id, &book.Name, &book.Author.Id)
		if err != nil {
			return domain.Book{}, err
		}
	}
	return book, nil
}

func (sql *bookPostgres) GetBooks() ([]domain.Book, error) {
	query := `SELECT id, name, author_id FROM books`
	rows, err := sql.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var book domain.Book
		err := rows.Scan(&book.Id, &book.Name, &book.Author.Id)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (sql *bookPostgres) UpdateBook(book domain.Book) error {
	query := `UPDATE books SET name = $1, author_id = $2 WHERE id = $3`
	_, err := sql.db.Exec(query, book.Name, book.Author.Id, book.Id)
	if err != nil {
		return err
	}
	return nil
}

func (sql *bookPostgres) DeleteBook(id int64) error {
	query := `DELETE FROM books WHERE id = $1`
	_, err := sql.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

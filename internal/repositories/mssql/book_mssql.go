package mssql

import (
	"database/sql"

	"github.com/hsnkorkmaz/go-hexagonal/internal/core/domain"
)

type bookMssql struct {
	db *sql.DB
}

func NewBookMssql(db *sql.DB) *bookMssql {
	return &bookMssql{db}
}

func (sql *bookMssql) CreateBook(book domain.Book) error {
	query := `INSERT INTO books (name, author_id) VALUES (?, ?)`
	_, err := sql.db.Exec(query, book.Name, book.Author.Id)
	if err != nil {
		return err
	}
	return nil
}

func (sql *bookMssql) GetBook(id int64) (domain.Book, error) {
	query := `SELECT id, name, author_id FROM books WHERE id = ?`
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

func (sql *bookMssql) GetBooks() ([]domain.Book, error) {
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

func (sql *bookMssql) UpdateBook(book domain.Book) error {
	query := `UPDATE books SET name = ?, author_id = ? WHERE id = ?`
	_, err := sql.db.Exec(query, book.Name, book.Author.Id, book.Id)
	if err != nil {
		return err
	}
	return nil
}

func (sql *bookMssql) DeleteBook(id int64) error {
	query := `DELETE FROM books WHERE id = ?`
	_, err := sql.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

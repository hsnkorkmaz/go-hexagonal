package mssql

import (
	"database/sql"
	"github.com/hsnkorkmaz/go-hexagonal/internal/core/domain"
)

type mssql struct {
	db *sql.DB
}

func NewMssql(db *sql.DB) *mssql {
	return &mssql{db}
}

//internal functions
func (sql *mssql) sqlQuery(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := sql.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (sql *mssql) sqlExec(query string, args ...interface{}) (sql.Result, error) {
	result, err := sql.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//books
func (sql *mssql) CreateBook(book domain.Book) error {
	query := `INSERT INTO books (name, author_id) VALUES (?, ?)`
	_, err := sql.sqlExec(query, book.Name, book.Author.Id)
	if err != nil {
		return err
	}
	return nil
}

func (sql *mssql) GetBook(id int64) (domain.Book, error) {
	query := `SELECT id, name, author_id FROM books WHERE id = ?`
	rows, err := sql.sqlQuery(query, id)
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

func (sql *mssql) GetBooks() ([]domain.Book, error) {
	query := `SELECT id, name, author_id FROM books`
	rows, err := sql.sqlQuery(query)
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

func (sql *mssql) UpdateBook(book domain.Book) error {
	query := `UPDATE books SET name = ?, author_id = ? WHERE id = ?`
	_, err := sql.sqlExec(query, book.Name, book.Author.Id, book.Id)
	if err != nil {
		return err
	}
	return nil
}

func (sql *mssql) DeleteBook(id int64) error {
	query := `DELETE FROM books WHERE id = ?`
	_, err := sql.sqlExec(query, id)
	if err != nil {
		return err
	}
	return nil
}

//authors
func (sql *mssql) CreateAuthor(author domain.Author) error {
	query := `INSERT INTO authors (name) VALUES (?)`
	_, err := sql.sqlExec(query, author.Name)
	if err != nil {
		return err
	}
	return nil
}

func (sql *mssql) GetAuthor(id int64) (domain.Author, error) {
	query := `SELECT id, name FROM authors WHERE id = ?`
	rows, err := sql.sqlQuery(query, id)
	if err != nil {
		return domain.Author{}, err
	}
	defer rows.Close()

	var author domain.Author
	for rows.Next() {
		err := rows.Scan(&author.Id, &author.Name)
		if err != nil {
			return domain.Author{}, err
		}
	}
	return author, nil
}

func (sql *mssql) GetAuthors() ([]domain.Author, error) {
	query := `SELECT id, name FROM authors`
	rows, err := sql.sqlQuery(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []domain.Author
	for rows.Next() {
		var author domain.Author
		err := rows.Scan(&author.Id, &author.Name)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (sql *mssql) UpdateAuthor(author domain.Author) error {
	query := `UPDATE authors SET name = ? WHERE id = ?`
	_, err := sql.sqlExec(query, author.Name, author.Id)
	if err != nil {
		return err
	}
	return nil
}

func (sql *mssql) DeleteAuthor(id int64) error {
	query := `DELETE FROM authors WHERE id = ?`
	_, err := sql.sqlExec(query, id)
	if err != nil {
		return err
	}
	return nil
}

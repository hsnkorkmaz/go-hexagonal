package mssql

import (
	"database/sql"

	"github.com/hsnkorkmaz/go-hexagonal/internal/core/domain"
)

type authorMssql struct {
	db *sql.DB
}

func NewAuthorMssql(db *sql.DB) *authorMssql {
	return &authorMssql{db}
}

func (sql *authorMssql) CreateAuthor(author domain.Author) error {
	query := `INSERT INTO authors (name) VALUES (?)`
	_, err := sql.db.Exec(query, author.Name)
	if err != nil {
		return err
	}
	return nil
}

func (sql *authorMssql) GetAuthor(id int64) (domain.Author, error) {
	query := `SELECT id, name FROM authors WHERE id = ?`
	rows, err := sql.db.Query(query, id)
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

func (sql *authorMssql) GetAuthors() ([]domain.Author, error) {
	query := `SELECT id, name FROM authors`
	rows, err := sql.db.Query(query)
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

func (sql *authorMssql) UpdateAuthor(author domain.Author) error {
	query := `UPDATE authors SET name = ? WHERE id = ?`
	_, err := sql.db.Exec(query, author.Name, author.Id)
	if err != nil {
		return err
	}
	return nil
}

func (sql *authorMssql) DeleteAuthor(id int64) error {
	query := `DELETE FROM authors WHERE id = ?`
	_, err := sql.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

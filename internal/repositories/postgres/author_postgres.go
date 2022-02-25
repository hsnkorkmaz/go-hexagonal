package postgres

import (
	"database/sql"

	"github.com/hsnkorkmaz/go-hexagonal/internal/core/domain"
	_ "github.com/lib/pq"
)

type authorPostgres struct {
	db *sql.DB
}

func NewAuthorPostgres(db *sql.DB) *authorPostgres {
	return &authorPostgres{db}
}

func (sql *authorPostgres) CreateAuthor(author domain.Author) error {
	query := `INSERT INTO authors (name) VALUES ($1)`
	_, err := sql.db.Exec(query, author.Name)
	if err != nil {
		return err
	}
	return nil
}

func (sql *authorPostgres) GetAuthor(id int64) (domain.Author, error) {
	query := `SELECT id, name FROM authors WHERE id = $1`
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

func (sql *authorPostgres) GetAuthors() ([]domain.Author, error) {
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

func (sql *authorPostgres) UpdateAuthor(author domain.Author) error {
	query := `UPDATE authors SET name = $1 WHERE id = $2`
	_, err := sql.db.Exec(query, author.Name, author.Id)
	if err != nil {
		return err
	}
	return nil
}

func (sql *authorPostgres) DeleteAuthor(id int64) error {
	query := `DELETE FROM authors WHERE id = $1`
	_, err := sql.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

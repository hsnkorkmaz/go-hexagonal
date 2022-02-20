package main

import (
	"database/sql"
	"fmt"
	//"hex-arch/internal/core/domain"
	"github.com/hsnkorkmaz/go-hexagonal/internal/core/services"
	"github.com/hsnkorkmaz/go-hexagonal/internal/repositories/mssql"
	_ "github.com/denisenkom/go-mssqldb"
)

const SQL_SERVER = ""
const SQL_PORT = ""
const SQL_USER = ""
const SQL_PASSWORD = ""
const SQL_DATABASE = ""

//cli
func main() {
	sqlCon, err := openSqlConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlRepository := mssql.NewMssql(sqlCon)
	bookService := services.NewBookService(sqlRepository)
	//authorService := services.NewAuthorService(sqlRepository)


	/* authorError := authorService.CreateAuthor(domain.Author{Name: "Hasan"})
	if authorError != nil {
		fmt.Println(authorError)
		return
	} */

	/* error := bookService.CreateBook(domain.Book{
		Name: "Book 2",
		Author: domain.Author{
			Id: 1,
		},
	})

	if error != nil {
		fmt.Println(error)
		return
	}


	error2 := bookService.CreateBook(domain.Book{
		Name: "Book 3",
		Author: domain.Author{
			Id: 1,
		},
	})

	if error2 != nil {
		fmt.Println(error2)
		return
	} */


	books, getError := bookService.GetBooks()

	if getError != nil {
		fmt.Println(getError)
		return
	}

	for _, book := range books {
		fmt.Println(book)
	}

}

func openSqlConnection() (*sql.DB, error) {
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		SQL_USER, SQL_PASSWORD, SQL_SERVER, SQL_PORT, SQL_DATABASE)

	SQL_DB, err := sql.Open("mssql", connectionString)
	if err != nil {
		return nil, err
	}
	return SQL_DB, nil
}

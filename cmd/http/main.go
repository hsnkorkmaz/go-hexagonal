package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"github.com/hsnkorkmaz/go-hexagonal/config/local"
	"github.com/hsnkorkmaz/go-hexagonal/internal/core/services"
	httpH "github.com/hsnkorkmaz/go-hexagonal/internal/handlers/http"
	"github.com/hsnkorkmaz/go-hexagonal/internal/repositories/mssql"
)

//http
func main() {
	// get sql connection
	sqlCon, err := openSqlConnection(local.SQL_USER, local.SQL_PASSWORD, local.SQL_SERVER, local.SQL_PORT, local.SQL_DATABASE)
	if err != nil {
		fmt.Println(err)
		return
	}

	sqlRepository := mssql.NewMssql(sqlCon)
	bookService := services.NewBookService(sqlRepository)
	authorService := services.NewAuthorService(sqlRepository)
	authorHandler := httpH.NewAuthorHandler(authorService)
	bookHandler := httpH.NewBookHandler(bookService)

	//http handlers
	router := mux.NewRouter()
	router.HandleFunc("/", nil).Methods("GET")

	//author handlers
	router.HandleFunc("/authors", authorHandler.GetAuthors).Methods("GET")
	router.HandleFunc("/authors", authorHandler.CreateAuthor).Methods("POST")
	router.HandleFunc("/authors/{id:[0-9]+}", authorHandler.GetAuthor).Methods("GET")
	router.HandleFunc("/authors/{id:[0-9]+}", authorHandler.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/authors/{id:[0-9]+}", authorHandler.DeleteAuthor).Methods("DELETE")

	// book handlers
	router.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")
	router.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id:[0-9]+}", bookHandler.GetBook).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}", bookHandler.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id:[0-9]+}", bookHandler.DeleteBook).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}

func openSqlConnection(sqlUser string, sqlPassword string, sqlServer string, sqlPort string, sqlDatabase string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		sqlUser, sqlPassword, sqlServer, sqlPort, sqlDatabase)

	SQL_DB, err := sql.Open("mssql", connectionString)
	if err != nil {
		return nil, err
	}
	return SQL_DB, nil
}

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
	"github.com/hsnkorkmaz/go-hexagonal/internal/middleware"
	"github.com/hsnkorkmaz/go-hexagonal/internal/repositories/mssql"
)

//http
func main() {

	// get sql connection
	sqlCon, err := openSqlConnection(local.SQL_USER, local.SQL_PASSWORD, local.SQL_SERVER, local.SQL_PORT, local.SQL_DATABASE)
	if err != nil {
		panic(err)
	}

	bookRepository := mssql.NewBookMssql(sqlCon)
	authorRepository := mssql.NewAuthorMssql(sqlCon)
	bookService := services.NewBookService(bookRepository)
	authorService := services.NewAuthorService(authorRepository)
	authorHandler := httpH.NewAuthorHandler(authorService)
	bookHandler := httpH.NewBookHandler(bookService)

	//auth middleware can work with Azure AD or Azure AD B2C
	//since public encription keys are different invoke corresponding function to get the keyset url
	//getADKeysUrl() or getADB2CKeysUrl(local.AZURE_AD_B2C_TENANT_NAME, local.AZURE_AD_B2C_USER_FLOW)
	accessControl := middleware.NewRBAC()
	authorisation := middleware.NewAzureMiddleware(getADKeysUrl(), accessControl)

	//http handlers
	router := mux.NewRouter()
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Api is running"))
	})

	//author handlers
	router.HandleFunc("/authors", authorisation.AzureMiddleWare(authorHandler.GetAuthors, []string{"Authors.Read"})).Methods("GET")
	router.HandleFunc("/authors", authorisation.AzureMiddleWare(authorHandler.CreateAuthor, []string{"Authors.Create"})).Methods("POST")
	router.HandleFunc("/authors/{id:[0-9]+}", authorisation.AzureMiddleWare(authorHandler.GetAuthor, []string{"Authors.Read"})).Methods("GET")
	router.HandleFunc("/authors/{id:[0-9]+}", authorisation.AzureMiddleWare(authorHandler.UpdateAuthor, []string{"Authors.Read", "Authors.Create"})).Methods("PUT")
	router.HandleFunc("/authors/{id:[0-9]+}", authorHandler.DeleteAuthor).Methods("DELETE")

	// book handlers
	router.HandleFunc("/books", authorisation.AzureMiddleWare(bookHandler.GetBooks, []string{"Books.Read"})).Methods("GET")
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

func getADB2CKeysUrl(tenant string, policy string) string {
	return "https://" + tenant + ".b2clogin.com/" + tenant + ".onmicrosoft.com/" + policy + "/discovery/v2.0/keys"
}

func getADKeysUrl() string {
	return "https://login.microsoftonline.com/common/discovery/v2.0/keys"
}

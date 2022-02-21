### Hexagonal Architecture Implementation in Go


Project Structure
```bash

│   .gitignore
│   go.mod
│   go.sum
│   main.go
│   README.md
│
├───cmd
│   ├───cli
│   │       main.go
│   │
│   └───http
│           main.go
│
├───config
│   ├───local
│   │       config.go
│   │
│   └───server
│           config.go
│
└───internal
    ├───adapters
    ├───core
    │   ├───domain
    │   │       author.go
    │   │       book.go
    │   │
    │   ├───ports
    │   │       Irepositories.go
    │   │       Iservices.go
    │   │
    │   └───services
    │           author_service.go
    │           book_service.go
    │
    ├───handlers
    │   ├───grpc
    │   └───http
    │           author_handler.go
    │           book_handler.go
    │
    └───repositories
        ├───mongo
        ├───mssql
        │       author_mssql.go
        │       book_mssql.go
        │
        └───postgres
```
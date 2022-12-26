# RestAPI with Gin Gonic

We will use RestAPI with Gin Gonic for this project.

## Requirement

1. You have PostgreSQL in your computer or you can use the cloud instead.
2. For this project, we will need some packages. The dependencies that you need:

```go
  go get -u github.com/gin-gonic/gin
  go get github.com/lib/pq
  go get github.com/joho/godotenv
```

- Gin is one of frameworks that most used in Go.
- PQ is package that required to `connect Go to PostgreSQL database`
- Godotenv is package to read the `.env` file

## How to use without docker

1. You can clone this project.
2. You have to make database `recordings` in the postgres. Run this to terminal:
   ```zsh
    psql -U postgres
    postgres=# CREATE database recordings;
    postgres=# exit
   ```
3. Then you can import this sql file to database with running this command:
   ```zsh
    psql -d recordings -f recordings.sql
   ```
4. Or you can run manually to create database, create table, and insert data with `DBeaver`, `PGAdmin`, and so on.
5. You can run:
   ```go
    go mod tidy
   ```
   to install the dependencies.
6. And finally, you can run with
   ```go
    go run main.go
   ```
7. You can just run the handler that do you want in `main.go`.
8. For documentations, you can import `postman-docs/Edspert.postman_collection.json` file to your Postman if you want to test the API.

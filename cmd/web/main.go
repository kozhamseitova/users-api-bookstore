package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	users "gitlab.com/tleuzhan13/bookstore/users-api/domain"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type app struct {
	router   *gin.Engine
	errorLog *log.Logger
	infoLog  *log.Logger
	users    *users.SnippetModel
}

func main() {

	addr := flag.String("addr", ":8082", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	pool, err := pgxpool.Connect(context.Background(), "user=admin password=123 host=localhost port=5432 dbname=snippetbox sslmode=disable pool_max_conns=10")
	if err != nil {
		log.Fatalf("Unable to connection to database: %v\n", err)
	}

	defer pool.Close()

	router := gin.Default()

	app := &app{
		router:   router,
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	app.routes()

	infoLog.Printf("Starting server on %s", *addr)
	err := router.Run(*addr)
	if err != nil {
		app.errorLog.Fatal(err)
	}
}
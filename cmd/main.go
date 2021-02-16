package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type app struct {
	router *gin.Engine
	errorLog      *log.Logger
	infoLog       *log.Logger

}

func main() {

	addr := flag.String("addr", ":8082", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)


	router := gin.Default()

	app := &app{
		router: router,
		errorLog: errorLog,
		infoLog: infoLog,
	}
	app.routes()

	infoLog.Printf("Starting server on %s", *addr)
	err := router.Run(*addr)
	if err != nil {
		app.errorLog.Fatal(err)
	}
}

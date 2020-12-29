//  Golang service API for HotelScanner
//
//  Swagger spec.
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//	- application/json
//  swagger:meta
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Kostikans/avitoTest/internal/package/logger"

	apiMiddleware "github.com/Kostikans/avitoTest/internal/app/middleware"

	"github.com/Kostikans/avitoTest/configs"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("../api/swagger")))

	return router
}

func InitDB() *sqlx.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		configs.BdConfig.User,
		configs.BdConfig.Password,
		configs.BdConfig.DBName)

	fmt.Println(connStr)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func main() {
	db := InitDB()
	logOutput, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer logOutput.Close()
	log := logger.NewLogger(logOutput)
	if err != nil {
		log.Error(err)
	}

	r := NewRouter()
	r.Use(apiMiddleware.NewPanicMiddleware())
	r.Use(apiMiddleware.LoggerMiddleware(log))

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Server started at port", ":8080")
}

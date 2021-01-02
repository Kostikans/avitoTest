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

	bookingDelivery "github.com/Kostikans/avitoTest/internal/app/booking/delivery/http"
	bookingRepository "github.com/Kostikans/avitoTest/internal/app/booking/repository"
	bookingUsecase "github.com/Kostikans/avitoTest/internal/app/booking/usecase"

	"github.com/joho/godotenv"

	roomDelivery "github.com/Kostikans/avitoTest/internal/app/room/delivery/http"
	roomRepository "github.com/Kostikans/avitoTest/internal/app/room/repository"
	roomUsecase "github.com/Kostikans/avitoTest/internal/app/room/usecase"

	"github.com/Kostikans/avitoTest/internal/package/logger"

	apiMiddleware "github.com/Kostikans/avitoTest/internal/app/middleware"

	"github.com/Kostikans/avitoTest/configs"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./api/swagger")))

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

	err := godotenv.Load("vars.env")
	if err != nil {
		log.Fatal(err)
	}
	configs.Init()
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

	roomRepo := roomRepository.NewRoomRepository(db)
	bookingRepo := bookingRepository.NewBookingRepository(db)

	roomUse := roomUsecase.NewRoomUsecase(roomRepo)
	bookingUse := bookingUsecase.NewRoomUsecase(bookingRepo, roomRepo)

	roomDelivery.NewRoomHandler(r, roomUse, log)
	bookingDelivery.NewBookingHandler(r, bookingUse, log)

	err = http.ListenAndServe(":9000", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Server started at port", ":9000")
}

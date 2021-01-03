//  Golang service API for Avito
//
//  Swagger spec.
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/x-www-form-urlencoded
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

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/schema"

	"github.com/spf13/viper"

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
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:9000/swagger.yaml"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./api/swagger")))

	return router
}

func InitDB() *sqlx.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		configs.BdConfig.User,
		configs.BdConfig.Password,
		configs.BdConfig.DBName)

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
	err = configs.ExportConfig()
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
	bookingUse := bookingUsecase.NewBookingUsecase(bookingRepo, roomRepo)
	decoder := schema.NewDecoder()
	roomDelivery.NewRoomHandler(r, roomUse, log, decoder)
	bookingDelivery.NewBookingHandler(r, bookingUse, log, decoder)

	err = http.ListenAndServe(viper.GetString(configs.ConfigFields.AvitoServicePort), r)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Server started at port", viper.GetString(configs.ConfigFields.AvitoServicePort))
}

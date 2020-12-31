package bookingDelivery

import (
	"net/http"
	"strconv"

	bookingModel "github.com/Kostikans/avitoTest/internal/app/booking/model"

	"github.com/Kostikans/avitoTest/internal/app/booking"
	"github.com/Kostikans/avitoTest/internal/package/clientError"
	"github.com/Kostikans/avitoTest/internal/package/customError"
	"github.com/Kostikans/avitoTest/internal/package/logger"
	"github.com/Kostikans/avitoTest/internal/package/responses"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

type BookingHandler struct {
	bookingUsecase booking.Usecase
	log            *logger.CustomLogger
}

func NewBookingHandler(r *mux.Router, usecase booking.Usecase, log *logger.CustomLogger) *BookingHandler {
	handler := BookingHandler{bookingUsecase: usecase, log: log}

	r.HandleFunc("/api/v1/booking", handler.AddBooking).Methods("POST")
	r.HandleFunc("/api/v1/booking/{booking_id:[0-9]+}", handler.DeleteBooking).Methods("DELETE")
	return &handler
}

func (rh *BookingHandler) AddBooking(w http.ResponseWriter, r *http.Request) {
	bookingAdd := bookingModel.BookingAdd{}
	err := easyjson.UnmarshalFromReader(r.Body, &bookingAdd)
	if err != nil {
		customError.PostError(w, r, rh.log, err, clientError.BadRequest)
		return
	}

	bookingID, err := rh.bookingUsecase.AddBooking(bookingAdd)
	if err != nil {
		customError.PostError(w, r, rh.log, err, nil)
		return
	}
	responses.SendDataResponse(w, bookingID)
}

func (rh *BookingHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	bookingIDVar := mux.Vars(r)["booking_id"]
	bookingID, err := strconv.ParseInt(bookingIDVar, 10, 64)
	if err != nil {
		customError.PostError(w, r, rh.log, err, clientError.BadRequest)
		return
	}

	err = rh.bookingUsecase.DeleteRoom(bookingID)
	if err != nil {
		customError.PostError(w, r, rh.log, err, nil)
		return
	}
	responses.SendOkResponse(w)
}

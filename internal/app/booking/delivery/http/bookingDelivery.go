package bookingDelivery

import (
	"net/http"
	"strconv"

	"github.com/Kostikans/avitoTest/internal/package/okCodes"

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

	r.HandleFunc("/bookings/create", handler.AddBooking).Methods("POST")
	r.Path("/bookings/list").Queries("room_id", "{room_id:[0-9]+}").HandlerFunc(handler.GetBookings).Methods("GET")
	r.Path("/bookings/delete").Queries("booking_id", "{booking_id:[0-9]+}").HandlerFunc(handler.DeleteBooking).Methods("DELETE")
	return &handler
}

// swagger:route POST /bookings/create bookings AddBooking
// responses:
//  201: bookingID
//  400: badrequest
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
	responses.SendDataResponse(w, okCodes.CreateResponse, bookingID)
}

// swagger:route DELETE /bookings/delete bookings DeleteBooking
// responses:
//  400: badrequest
//  404: notfound
func (rh *BookingHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	bookingIDVar := r.FormValue("booking_id")
	bookingID, err := strconv.Atoi(bookingIDVar)
	if err != nil {
		customError.PostError(w, r, rh.log, err, clientError.BadRequest)
		return
	}

	err = rh.bookingUsecase.DeleteBooking(bookingID)
	if err != nil {
		customError.PostError(w, r, rh.log, err, nil)
		return
	}
	responses.SendOkResponse(w, okCodes.OkResponse)
}

// swagger:route GET /bookings/list bookings GetBookings
// responses:
//  200: bookings
//  400: badrequest
//  404: notfound
func (rh *BookingHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	roomIDVar := r.FormValue("room_id")
	roomID, err := strconv.Atoi(roomIDVar)
	if err != nil {
		customError.PostError(w, r, rh.log, err, clientError.BadRequest)
		return
	}
	rooms, err := rh.bookingUsecase.GetBookings(roomID)
	if err != nil {
		customError.PostError(w, r, rh.log, err, nil)
		return
	}
	responses.SendDataResponse(w, okCodes.OkResponse, rooms)
}

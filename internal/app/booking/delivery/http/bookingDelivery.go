package bookingDelivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/schema"

	"github.com/Kostikans/avitoTest/internal/package/okCodes"

	bookingModel "github.com/Kostikans/avitoTest/internal/app/booking/model"

	"github.com/Kostikans/avitoTest/internal/app/booking"
	"github.com/Kostikans/avitoTest/internal/package/clientError"
	"github.com/Kostikans/avitoTest/internal/package/customError"
	"github.com/Kostikans/avitoTest/internal/package/logger"
	"github.com/Kostikans/avitoTest/internal/package/responses"
	"github.com/gorilla/mux"
)

type BookingHandler struct {
	bookingUsecase booking.Usecase
	decoder        *schema.Decoder
	log            *logger.CustomLogger
}

func NewBookingHandler(r *mux.Router, usecase booking.Usecase, log *logger.CustomLogger, decoder *schema.Decoder) *BookingHandler {
	handler := BookingHandler{bookingUsecase: usecase, log: log, decoder: decoder}

	r.HandleFunc("/bookings/create", handler.AddBooking).Methods("POST")
	r.Path("/bookings/list").Queries("room_id", "{room_id}").HandlerFunc(handler.GetBookings).Methods("GET")
	r.Path("/bookings/delete").Queries("booking_id", "{booking_id}").HandlerFunc(handler.DeleteBooking).Methods("DELETE")
	return &handler
}

// swagger:route POST /bookings/create bookings AddBooking
// responses:
//  200:
//  201: bookingID
//  400: badrequest
func (bh *BookingHandler) AddBooking(w http.ResponseWriter, r *http.Request) {
	bookingAdd := bookingModel.BookingAdd{}
	err := r.ParseForm()
	if err != nil {
		customError.PostError(w, r, bh.log, errors.New("cannot parse form"), clientError.BadRequest)
		return
	}
	err = bh.decoder.Decode(&bookingAdd, r.PostForm)
	if err != nil {
		customError.PostError(w, r, bh.log, errors.New("cannot decode into struct"), clientError.BadRequest)
		return
	}
	bookingID, err := bh.bookingUsecase.AddBooking(bookingAdd)
	if err != nil {
		customError.PostError(w, r, bh.log, err, nil)
		return
	}
	responses.SendDataResponse(w, okCodes.CreateResponse, bookingID)
}

// swagger:route DELETE /bookings/delete bookings DeleteBooking
// responses:
//  200:
//  400: badrequest
//  404: notfound
func (bh *BookingHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	bookingIDVar := r.FormValue("booking_id")
	bookingID, err := strconv.Atoi(bookingIDVar)
	if err != nil {
		customError.PostError(w, r, bh.log, errors.New("wrong type of query params"), clientError.BadRequest)
		return
	}

	err = bh.bookingUsecase.DeleteBooking(bookingID)
	if err != nil {
		customError.PostError(w, r, bh.log, err, nil)
		return
	}
	responses.SendOkResponse(w, okCodes.OkResponse)
}

// swagger:route GET /bookings/list bookings GetBookings
// responses:
//  200: bookings
//  400: badrequest
//  404: notfound
func (bh *BookingHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	roomIDVar := r.FormValue("room_id")
	roomID, err := strconv.Atoi(roomIDVar)
	if err != nil {
		customError.PostError(w, r, bh.log, errors.New("wrong type of query params"), clientError.BadRequest)
		return
	}
	rooms, err := bh.bookingUsecase.GetBookings(roomID)
	if err != nil {
		customError.PostError(w, r, bh.log, err, nil)
		return
	}
	responses.SendDataResponse(w, okCodes.OkResponse, rooms)
}

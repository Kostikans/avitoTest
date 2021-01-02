package bookingRepository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Kostikans/avitoTest/internal/package/clientError"

	bookingModel "github.com/Kostikans/avitoTest/internal/app/booking/model"
	"github.com/Kostikans/avitoTest/internal/package/serverError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/jmoiron/sqlx"
)

type BookingRepository struct {
	db *sqlx.DB
}

func NewBookingRepository(db *sqlx.DB) *BookingRepository {
	return &BookingRepository{db: db}
}

func (rRep *BookingRepository) AddBooking(booking bookingModel.BookingAdd) (bookingModel.BookingID, error) {
	bookingID := bookingModel.BookingID{}
	dateStart, err := time.Parse("2006-01-02", booking.DateStart)
	if err != nil {
		return bookingID, customerror.NewCustomError(errors.New("error while parse input time"), clientError.BadRequest, 1)
	}
	dateEnd, err := time.Parse("2006-01-02", booking.DateEnd)
	if err != nil {
		return bookingID, customerror.NewCustomError(errors.New("error while parse input time"), clientError.BadRequest, 1)
	}
	if dateStart.After(dateEnd) {
		return bookingID, customerror.NewCustomError(errors.New("dateStart stands after dateEnd"), clientError.BadRequest, 1)
	}
	err = rRep.db.QueryRow(AddBookingPostgreRequest, booking.RoomID, dateStart,
		dateEnd).Scan(&bookingID.BookingID)
	if err != nil {
		return bookingID, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return bookingID, nil
}

func (rRep *BookingRepository) DeleteBooking(bookingID int) error {
	_, err := rRep.db.Exec(DeleteRoomPostgreRequest, bookingID)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return err
}

func (rRep *BookingRepository) GetBookings(roomID int) ([]bookingModel.Booking, error) {
	var bookings []bookingModel.Booking
	err := rRep.db.Select(&bookings, GetBookingsPostgreRequest, roomID)
	if err != nil {
		return bookings, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return bookings, err
}

func (rRep *BookingRepository) CheckBookingExist(roomID int) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT exists (%s)", CheckBookingExistPostgreRequest)
	err := rRep.db.QueryRow(query, roomID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return exists, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return exists, nil
}

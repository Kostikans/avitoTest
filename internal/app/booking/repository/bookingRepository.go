package bookingRepository

import (
	"time"

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
	dateStart, _ := time.Parse("HHHH-MM-HH", booking.DateStart)
	dateEnd, _ := time.Parse("HHHH-MM-HH", booking.DateEnd)
	err := rRep.db.QueryRow(AddBookingPostgreRequest, booking.RoomID, dateStart,
		dateEnd).Scan(&bookingID.BookingID)
	if err != nil {
		return bookingID, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return bookingID, nil
}

func (rRep *BookingRepository) DeleteRoom(bookingID int64) error {
	_, err := rRep.db.Exec(DeleteRoomPostgreRequest, bookingID)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return err
}

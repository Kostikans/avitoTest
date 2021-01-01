package booking

import bookingModel "github.com/Kostikans/avitoTest/internal/app/booking/model"

type Usecase interface {
	AddBooking(booking bookingModel.BookingAdd) (bookingModel.BookingID, error)
	DeleteBooking(bookingID int64) error
	GetBookings(roomID int64) ([]bookingModel.Booking, error)
	CheckBookingExist(bookingID int64) (bool, error)
}

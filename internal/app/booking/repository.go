//go:generate mockgen -source repository.go -destination mocks/booking_repository_mock.go -package booking_mock
package booking

import bookingModel "github.com/Kostikans/avitoTest/internal/app/booking/model"

type Repository interface {
	AddBooking(booking bookingModel.BookingAdd) (bookingModel.BookingID, error)
	DeleteBooking(bookingID int) error
	GetBookings(roomID int) ([]bookingModel.Booking, error)
	CheckBookingExist(bookingID int) (bool, error)
}

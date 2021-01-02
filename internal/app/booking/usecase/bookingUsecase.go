package bookingUsecase

import (
	"errors"

	"github.com/Kostikans/avitoTest/internal/app/booking"
	bookingModel "github.com/Kostikans/avitoTest/internal/app/booking/model"
	"github.com/Kostikans/avitoTest/internal/app/room"
	"github.com/Kostikans/avitoTest/internal/package/clientError"
	"github.com/Kostikans/avitoTest/internal/package/customError"
)

type BookingUsecase struct {
	BookingRepo booking.Repository
	RoomRepo    room.Repository
}

func NewRoomUsecase(BookingRepo booking.Repository, RoomRepo room.Repository) *BookingUsecase {
	return &BookingUsecase{BookingRepo: BookingRepo, RoomRepo: RoomRepo}
}

func (bUsecase *BookingUsecase) AddBooking(booking bookingModel.BookingAdd) (bookingModel.BookingID, error) {
	bookingID := bookingModel.BookingID{}
	exist, err := bUsecase.RoomRepo.CheckRoomExist(booking.RoomID)
	if err != nil {
		return bookingID, err
	}
	if exist == false {
		return bookingID, customError.NewCustomError(errors.New("room doesn't exist"), clientError.NotFound, 1)
	}
	bookingID, err = bUsecase.BookingRepo.AddBooking(booking)
	return bookingID, err
}

func (bUsecase *BookingUsecase) DeleteBooking(bookingID int) error {
	exist, err := bUsecase.CheckBookingExist(bookingID)
	if err != nil {
		return err
	}
	if exist == false {
		return customError.NewCustomError(errors.New("booking doesn't exist"), clientError.NotFound, 1)
	}
	return bUsecase.BookingRepo.DeleteBooking(bookingID)
}

func (bUsecase *BookingUsecase) GetBookings(roomID int) ([]bookingModel.Booking, error) {
	var bookings []bookingModel.Booking
	exist, err := bUsecase.RoomRepo.CheckRoomExist(roomID)
	if err != nil {
		return bookings, err
	}
	if exist == false {
		return bookings, customError.NewCustomError(errors.New("room doesn't exist"), clientError.NotFound, 1)
	}
	bookings, err = bUsecase.BookingRepo.GetBookings(roomID)
	return bookings, err
}

func (bUsecase *BookingUsecase) CheckBookingExist(bookingID int) (bool, error) {
	return bUsecase.BookingRepo.CheckBookingExist(bookingID)
}

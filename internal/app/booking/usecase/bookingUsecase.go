package bookingUsecase

import (
	"github.com/Kostikans/avitoTest/internal/app/booking"
	bookingModel "github.com/Kostikans/avitoTest/internal/app/booking/model"
)

type BookingUsecase struct {
	BookingRepo booking.Repository
}

func NewRoomUsecase(BookingRepo booking.Repository) *BookingUsecase {
	return &BookingUsecase{BookingRepo: BookingRepo}
}

func (bUsecase *BookingUsecase) AddBooking(booking bookingModel.BookingAdd) (bookingModel.BookingID, error) {
	return bUsecase.BookingRepo.AddBooking(booking)
}

func (bUsecase *BookingUsecase) DeleteRoom(bookingID int64) error {
	return bUsecase.BookingRepo.DeleteRoom(bookingID)
}

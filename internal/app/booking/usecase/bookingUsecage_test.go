package bookingUsecase

import (
	"errors"
	"testing"

	"github.com/Kostikans/avitoTest/internal/package/serverError"

	"github.com/Kostikans/avitoTest/internal/package/clientError"
	"github.com/Kostikans/avitoTest/internal/package/customError"

	booking_mock "github.com/Kostikans/avitoTest/internal/app/booking/mocks"

	bookingModel "github.com/Kostikans/avitoTest/internal/app/booking/model"

	room_mock "github.com/Kostikans/avitoTest/internal/app/room/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBookingUsecase_AddBooking(t *testing.T) {
	t.Run("AddBooking", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		bookingAdd := bookingModel.BookingAdd{RoomID: 1, DateStart: "2013-10-12", DateEnd: "2013-10-12"}
		bookingIDTest := bookingModel.BookingID{BookingID: 1}
		roomID := 1
		mockRoomRepository := room_mock.NewMockRepository(ctrl)
		mockBookingRepository := booking_mock.NewMockRepository(ctrl)

		mockRoomRepository.EXPECT().
			CheckRoomExist(roomID).
			Return(true, nil)
		mockBookingRepository.EXPECT().
			AddBooking(bookingAdd).
			Return(bookingIDTest, nil)

		bookingUsecase := NewBookingUsecase(mockBookingRepository, mockRoomRepository)

		bookingID, err := bookingUsecase.AddBooking(bookingAdd)
		assert.NoError(t, err)
		assert.Equal(t, bookingID, bookingIDTest)
	})
	t.Run("AddBookingErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		bookingAdd := bookingModel.BookingAdd{RoomID: 1, DateStart: "2013-10-12", DateEnd: "2013-10-12"}
		roomID := 1
		mockRoomRepository := room_mock.NewMockRepository(ctrl)
		mockBookingRepository := booking_mock.NewMockRepository(ctrl)

		mockRoomRepository.EXPECT().
			CheckRoomExist(roomID).
			Return(false, nil)

		bookingUsecase := NewBookingUsecase(mockBookingRepository, mockRoomRepository)

		_, err := bookingUsecase.AddBooking(bookingAdd)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), clientError.NotFound)
	})
	t.Run("AddBookingErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		bookingAdd := bookingModel.BookingAdd{RoomID: 1, DateStart: "2013-10-12", DateEnd: "2013-10-12"}
		roomID := 1
		mockRoomRepository := room_mock.NewMockRepository(ctrl)
		mockBookingRepository := booking_mock.NewMockRepository(ctrl)

		mockRoomRepository.EXPECT().
			CheckRoomExist(roomID).
			Return(true, customError.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		bookingUsecase := NewBookingUsecase(mockBookingRepository, mockRoomRepository)

		_, err := bookingUsecase.AddBooking(bookingAdd)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestBookingUsecase_DeleteBooking(t *testing.T) {
	t.Run("DeleteBooking", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		bookingID := 1
		mockRoomRepository := room_mock.NewMockRepository(ctrl)
		mockBookingRepository := booking_mock.NewMockRepository(ctrl)

		mockBookingRepository.EXPECT().
			CheckBookingExist(bookingID).
			Return(true, nil)
		mockBookingRepository.EXPECT().
			DeleteBooking(bookingID).
			Return(nil)

		bookingUsecase := NewBookingUsecase(mockBookingRepository, mockRoomRepository)

		err := bookingUsecase.DeleteBooking(bookingID)
		assert.NoError(t, err)

	})
	t.Run("DeleteBookingErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		bookingID := 1
		mockRoomRepository := room_mock.NewMockRepository(ctrl)
		mockBookingRepository := booking_mock.NewMockRepository(ctrl)

		mockBookingRepository.EXPECT().
			CheckBookingExist(bookingID).
			Return(false, nil)

		bookingUsecase := NewBookingUsecase(mockBookingRepository, mockRoomRepository)

		err := bookingUsecase.DeleteBooking(bookingID)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), clientError.NotFound)
	})
	t.Run("DeleteBookingErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		bookingID := 1
		mockRoomRepository := room_mock.NewMockRepository(ctrl)
		mockBookingRepository := booking_mock.NewMockRepository(ctrl)

		mockBookingRepository.EXPECT().
			CheckBookingExist(bookingID).
			Return(false, customError.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		bookingUsecase := NewBookingUsecase(mockBookingRepository, mockRoomRepository)

		err := bookingUsecase.DeleteBooking(bookingID)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), serverError.ServerInternalError)

	})
}

func TestBookingUsecase_GetBookings(t *testing.T) {
	t.Run("GetBookings", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		bookingsTest := []bookingModel.Booking{
			{BookingID: 1, DateStart: "2013-11-10", DateEnd: "2014-12-12"},
			{BookingID: 2, DateStart: "2013-11-10", DateEnd: "2014-12-12"},
		}
		roomID := 1
		mockRoomRepository := room_mock.NewMockRepository(ctrl)
		mockBookingRepository := booking_mock.NewMockRepository(ctrl)

		mockRoomRepository.EXPECT().
			CheckRoomExist(roomID).
			Return(true, nil)
		mockBookingRepository.EXPECT().
			GetBookings(roomID).
			Return(bookingsTest, nil)

		bookingUsecase := NewBookingUsecase(mockBookingRepository, mockRoomRepository)

		bookings, err := bookingUsecase.GetBookings(roomID)
		assert.NoError(t, err)
		assert.Equal(t, bookings, bookingsTest)
	})
	t.Run("GetBookingsErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		roomID := 1
		mockRoomRepository := room_mock.NewMockRepository(ctrl)
		mockBookingRepository := booking_mock.NewMockRepository(ctrl)

		mockRoomRepository.EXPECT().
			CheckRoomExist(roomID).
			Return(false, nil)

		bookingUsecase := NewBookingUsecase(mockBookingRepository, mockRoomRepository)

		_, err := bookingUsecase.GetBookings(roomID)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), clientError.NotFound)
	})
	t.Run("GetBookingsErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		roomID := 1
		mockRoomRepository := room_mock.NewMockRepository(ctrl)
		mockBookingRepository := booking_mock.NewMockRepository(ctrl)

		mockRoomRepository.EXPECT().
			CheckRoomExist(roomID).
			Return(true, customError.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		bookingUsecase := NewBookingUsecase(mockBookingRepository, mockRoomRepository)

		_, err := bookingUsecase.GetBookings(roomID)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), serverError.ServerInternalError)
	})
}

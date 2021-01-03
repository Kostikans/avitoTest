package bookingDelivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gorilla/schema"

	bookingModel "github.com/Kostikans/avitoTest/internal/app/booking/model"

	booking_mock "github.com/Kostikans/avitoTest/internal/app/booking/mocks"

	"github.com/Kostikans/avitoTest/internal/package/clientError"
	"github.com/Kostikans/avitoTest/internal/package/customError"
	"github.com/Kostikans/avitoTest/internal/package/logger"
	"github.com/Kostikans/avitoTest/internal/package/okCodes"
	"github.com/Kostikans/avitoTest/internal/package/responses"
	"github.com/Kostikans/avitoTest/internal/package/serverError"
	"github.com/golang/mock/gomock"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestBookingHandler_AddBooking(t *testing.T) {
	decoder := schema.NewDecoder()
	t.Run("AddBooking", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testBookingAdd := bookingModel.BookingAdd{RoomID: 2, DateStart: "2013-08-23", DateEnd: "2014-09-23"}
		bookingIDTest := bookingModel.BookingID{BookingID: 2}
		mockBookingUseCase := booking_mock.NewMockUsecase(ctrl)

		var encoder = schema.NewEncoder()
		form := url.Values{}
		err := encoder.Encode(testBookingAdd, form)
		assert.NoError(t, err)

		mockBookingUseCase.EXPECT().
			AddBooking(testBookingAdd).
			Return(bookingIDTest, nil)

		req, err := http.NewRequest("GET", "bookings/create", nil)

		assert.NoError(t, err)
		req.PostForm = form

		rec := httptest.NewRecorder()
		handler := BookingHandler{
			bookingUsecase: mockBookingUseCase,
			log:            logger.NewLogger(os.Stdout),
			decoder:        decoder,
		}

		handler.AddBooking(rec, req)
		resp := rec.Result()

		bookingID := bookingModel.BookingID{}

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &bookingID)
		assert.NoError(t, err)

		assert.Equal(t, bookingIDTest, bookingID)
		assert.Equal(t, okCodes.CreateResponse, response.Code)
	})
	t.Run("AddBookingErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testBookingAdd := "fdscx:fdsf"

		mockBookingUseCase := booking_mock.NewMockUsecase(ctrl)

		body, err := json.Marshal(testBookingAdd)
		assert.NoError(t, err)

		req, err := http.NewRequest("GET", "bookings/create", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := BookingHandler{
			bookingUsecase: mockBookingUseCase,
			log:            logger.NewLogger(os.Stdout),
			decoder:        decoder,
		}

		handler.AddBooking(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})
	t.Run("AddBooking", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testBookingAdd := bookingModel.BookingAdd{RoomID: 2, DateStart: "2013-08-23", DateEnd: "2014-09-23"}
		bookingIDTest := bookingModel.BookingID{BookingID: 2}
		mockBookingUseCase := booking_mock.NewMockUsecase(ctrl)

		var encoder = schema.NewEncoder()
		form := url.Values{}
		err := encoder.Encode(testBookingAdd, form)
		assert.NoError(t, err)

		mockBookingUseCase.EXPECT().
			AddBooking(testBookingAdd).
			Return(bookingIDTest, customError.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		req, err := http.NewRequest("GET", "bookings/create", nil)
		assert.NoError(t, err)
		req.PostForm = form

		rec := httptest.NewRecorder()
		handler := BookingHandler{
			bookingUsecase: mockBookingUseCase,
			log:            logger.NewLogger(os.Stdout),
			decoder:        decoder,
		}

		handler.AddBooking(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)

		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})
}

func TestBookingHandler_GetBookings(t *testing.T) {
	t.Run("GetBooking", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testGetBookings := []bookingModel.Booking{
			{BookingID: 2, DateStart: "2013-12-23", DateEnd: "2014-11-23"},
			{BookingID: 3, DateStart: "2013-12-23", DateEnd: "2014-11-23"},
		}
		mockBookingUseCase := booking_mock.NewMockUsecase(ctrl)
		roomID := 2
		mockBookingUseCase.EXPECT().
			GetBookings(roomID).
			Return(testGetBookings, nil)

		req, err := http.NewRequest("GET", "bookings/list?room_id=2", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := BookingHandler{
			bookingUsecase: mockBookingUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.GetBookings(rec, req)
		resp := rec.Result()

		var GetBookings []bookingModel.Booking

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.([]interface{}), &GetBookings)
		assert.NoError(t, err)

		assert.Equal(t, testGetBookings, GetBookings)
		assert.Equal(t, okCodes.OkResponse, response.Code)
	})
	t.Run("GetBookingErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testGetBookings := []bookingModel.Booking{
			{BookingID: 2, DateStart: "2013-12-23", DateEnd: "2014-11-23"},
			{BookingID: 3, DateStart: "2013-12-23", DateEnd: "2014-11-23"},
		}
		mockBookingUseCase := booking_mock.NewMockUsecase(ctrl)
		roomID := 2
		mockBookingUseCase.EXPECT().
			GetBookings(roomID).
			Return(testGetBookings, customError.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		req, err := http.NewRequest("GET", "bookings/list?room_id=2", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := BookingHandler{
			bookingUsecase: mockBookingUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.GetBookings(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})
}

func TestBookingHandler_DeleteBooking(t *testing.T) {
	t.Run("DeleteBooking", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockBookingUseCase := booking_mock.NewMockUsecase(ctrl)
		bookingID := 2
		mockBookingUseCase.EXPECT().
			DeleteBooking(bookingID).
			Return(nil)

		req, err := http.NewRequest("GET", "bookings/delete?booking_id=2", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := BookingHandler{
			bookingUsecase: mockBookingUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.DeleteBooking(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)

		assert.Equal(t, okCodes.OkResponse, response.Code)
	})
	t.Run("DeleteBookingErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockBookingUseCase := booking_mock.NewMockUsecase(ctrl)
		bookingID := 2
		mockBookingUseCase.EXPECT().
			DeleteBooking(bookingID).
			Return(customError.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		req, err := http.NewRequest("GET", "bookings/delete?booking_id=2", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := BookingHandler{
			bookingUsecase: mockBookingUseCase,
			log:            logger.NewLogger(os.Stdout),
		}

		handler.DeleteBooking(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)

		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})
}

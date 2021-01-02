package bookingRepository

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Kostikans/avitoTest/internal/package/clientError"

	"github.com/Kostikans/avitoTest/internal/package/customError"
	"github.com/Kostikans/avitoTest/internal/package/serverError"

	"github.com/DATA-DOG/go-sqlmock"
	bookingModel "github.com/Kostikans/avitoTest/internal/app/booking/model"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestBookingRepository_AddBooking(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("AddBooking", func(t *testing.T) {
		bookingID := bookingModel.BookingID{BookingID: 1}
		rowsBooking := sqlmock.NewRows([]string{"booking_id"}).AddRow(
			1)

		query := AddBookingPostgreRequest

		bookingTest := bookingModel.BookingAdd{RoomID: 1, DateStart: "2010-01-30", DateEnd: "2013-01-24"}
		mock.ExpectQuery(query).
			WithArgs(bookingTest.RoomID, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(rowsBooking)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewBookingRepository(sqlxDb)

		bookingIDTest, err := rep.AddBooking(bookingTest)
		assert.NoError(t, err)
		assert.Equal(t, bookingID, bookingIDTest)
	})

	t.Run("AddBookingErr1", func(t *testing.T) {
		query := AddBookingPostgreRequest

		bookingTest := bookingModel.BookingAdd{RoomID: 1, DateStart: "2010-01-30", DateEnd: "2013-01-24"}
		mock.ExpectQuery(query).
			WithArgs(bookingTest.RoomID, bookingTest.DateStart, bookingTest.DateEnd).
			WillReturnError(customError.NewCustomError(errors.New("err"), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewBookingRepository(sqlxDb)

		_, err := rep.AddBooking(bookingTest)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), serverError.ServerInternalError)
	})

	t.Run("AddBookingErr2", func(t *testing.T) {
		bookingTest := bookingModel.BookingAdd{RoomID: 1, DateStart: "20101-01-30", DateEnd: "2013-01-24"}

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewBookingRepository(sqlxDb)

		_, err := rep.AddBooking(bookingTest)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), clientError.BadRequest)
	})

	t.Run("AddBookingErr2", func(t *testing.T) {
		bookingTest := bookingModel.BookingAdd{RoomID: 1, DateStart: "2010-01-30", DateEnd: "2013-01-244"}

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewBookingRepository(sqlxDb)

		_, err := rep.AddBooking(bookingTest)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), clientError.BadRequest)
	})

	t.Run("AddBookingErr3", func(t *testing.T) {
		bookingTest := bookingModel.BookingAdd{RoomID: 1, DateStart: "2014-01-30", DateEnd: "2013-01-24"}

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewBookingRepository(sqlxDb)

		_, err := rep.AddBooking(bookingTest)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), clientError.BadRequest)
	})
}

func TestBookingRepository_DeleteBooking(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("DeleteBooking", func(t *testing.T) {
		bookingID := 1

		query := DeleteRoomPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(bookingID).
			WillReturnRows()

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewBookingRepository(sqlxDb)

		err := rep.DeleteBooking(bookingID)
		assert.NoError(t, err)
	})
	t.Run("DeleteBookingErr", func(t *testing.T) {
		bookingID := 1

		query := DeleteRoomPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(bookingID).
			WillReturnError(customError.NewCustomError(errors.New("fds"), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewBookingRepository(sqlxDb)

		err := rep.DeleteBooking(bookingID)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestBookingRepository_CheckBookingExist(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("CheckBookingExist", func(t *testing.T) {
		query := fmt.Sprintf("SELECT exists (%s)", CheckBookingExistPostgreRequest)
		rowsBooking := sqlmock.NewRows([]string{"exists"}).AddRow(
			true)
		bookingID := 1
		mock.ExpectQuery(query).
			WithArgs(bookingID).
			WillReturnRows(rowsBooking)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewBookingRepository(sqlxDb)

		res, err := rep.CheckBookingExist(bookingID)
		assert.NoError(t, err)
		assert.Equal(t, res, true)
	})

	t.Run("CheckBookingExistErr", func(t *testing.T) {
		query := fmt.Sprintf("SELECT exists (%s)", CheckBookingExistPostgreRequest)
		bookingID := 1
		mock.ExpectQuery(query).
			WithArgs(bookingID).
			WillReturnError(customError.NewCustomError(errors.New("fds"), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewBookingRepository(sqlxDb)

		_, err := rep.CheckBookingExist(bookingID)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestBookingRepository_GetBookings(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("GetBookings", func(t *testing.T) {
		RoomID := 1
		rowsBooking := sqlmock.NewRows([]string{"booking_id", "date_start", "date_end"}).AddRow(
			1, "2013-11-10", "2014-12-12").AddRow(
			2, "2013-11-10", "2014-12-12")
		bookingsTest := []bookingModel.Booking{
			{BookingID: 1, DateStart: "2013-11-10", DateEnd: "2014-12-12"},
			{BookingID: 2, DateStart: "2013-11-10", DateEnd: "2014-12-12"},
		}
		query := GetBookingsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(RoomID).
			WillReturnRows(rowsBooking)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewBookingRepository(sqlxDb)

		bookings, err := rep.GetBookings(RoomID)
		assert.NoError(t, err)
		assert.Equal(t, bookings, bookingsTest)
	})
	t.Run("GetBookingsErr", func(t *testing.T) {
		RoomID := 1

		query := GetBookingsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs(RoomID).
			WillReturnError(customError.NewCustomError(errors.New("fds"), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewBookingRepository(sqlxDb)

		err := rep.DeleteBooking(RoomID)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), serverError.ServerInternalError)
	})
}

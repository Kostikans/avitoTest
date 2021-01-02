package roomRepository

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Kostikans/avitoTest/internal/package/customError"
	"github.com/Kostikans/avitoTest/internal/package/serverError"

	"github.com/DATA-DOG/go-sqlmock"
	roomModel "github.com/Kostikans/avitoTest/internal/app/room/model"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRoomRepository_AddRoom(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("AddRoom", func(t *testing.T) {
		roomID := roomModel.RoomID{RoomID: 1}
		rowsRoom := sqlmock.NewRows([]string{"room_id"}).AddRow(
			1)

		query := AddRoomPostgreRequest

		roomTest := roomModel.RoomAdd{Cost: 324, Description: "best room in the world"}
		mock.ExpectQuery(query).
			WithArgs(roomTest.Description, roomTest.Cost).
			WillReturnRows(rowsRoom)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewRoomRepository(sqlxDb)

		roomIDtest, err := rep.AddRoom(roomTest)
		assert.NoError(t, err)
		assert.Equal(t, roomID, roomIDtest)
	})

	t.Run("AddRoom", func(t *testing.T) {

		query := AddRoomPostgreRequest

		roomTest := roomModel.RoomAdd{Cost: 324, Description: "best room in the world"}
		mock.ExpectQuery(query).
			WithArgs(roomTest.Cost, roomTest.Description).
			WillReturnError(customError.NewCustomError(errors.New("fds"), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewRoomRepository(sqlxDb)

		_, err := rep.AddRoom(roomTest)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestRoomRepository_DeleteRoom(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("DeleteRoom", func(t *testing.T) {
		query := DeleteRoomPostgreRequest
		roomID := 1
		mock.ExpectQuery(query).
			WithArgs(roomID).
			WillReturnRows()

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewRoomRepository(sqlxDb)

		err = rep.DeleteRoom(roomID)
		assert.NoError(t, err)
	})

	t.Run("DeleteRoomErr", func(t *testing.T) {
		query := DeleteRoomPostgreRequest
		roomID := 1
		mock.ExpectQuery(query).
			WithArgs(roomID).
			WillReturnError(customError.NewCustomError(errors.New("fds"), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewRoomRepository(sqlxDb)

		err = rep.DeleteRoom(roomID)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestRoomRepository_GetRooms(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("GetRooms", func(t *testing.T) {
		roomTest := []roomModel.Room{
			{RoomID: 1, Cost: 234, Description: "nice room", Created: "2013-10-02"},
			{RoomID: 2, Cost: 324, Description: "best room", Created: "2014-20-03"},
		}
		rowsRoom := sqlmock.NewRows([]string{"room_id", "cost", "description", "created"}).AddRow(
			1, 234, "nice room", "2013-10-02").AddRow(2, 324, "best room", "2014-20-03")

		query := GetRoomsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(rowsRoom)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewRoomRepository(sqlxDb)
		roomOrder := roomModel.RoomOrder{}
		rooms, err := rep.GetRooms(roomOrder)
		assert.NoError(t, err)
		assert.Equal(t, roomTest, rooms)
	})

	t.Run("GetRooms1", func(t *testing.T) {
		roomTest := []roomModel.Room{
			{RoomID: 1, Cost: 234, Description: "nice room", Created: "2013-10-02"},
			{RoomID: 2, Cost: 324, Description: "best room", Created: "2014-20-03"},
		}
		rowsRoom := sqlmock.NewRows([]string{"room_id", "cost", "description", "created"}).AddRow(
			1, 234, "nice room", "2013-10-02").AddRow(2, 324, "best room", "2014-20-03")

		query := GetRoomsPostgreRequest
		query += " ORDER BY created "
		query += "DESC"

		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(rowsRoom)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewRoomRepository(sqlxDb)
		roomOrder := roomModel.RoomOrder{Sort: "date", Order: "true"}
		rooms, err := rep.GetRooms(roomOrder)
		assert.NoError(t, err)
		assert.Equal(t, roomTest, rooms)
	})
	t.Run("GetRooms2", func(t *testing.T) {
		roomTest := []roomModel.Room{
			{RoomID: 1, Cost: 234, Description: "nice room", Created: "2013-10-02"},
			{RoomID: 2, Cost: 324, Description: "best room", Created: "2014-20-03"},
		}
		rowsRoom := sqlmock.NewRows([]string{"room_id", "cost", "description", "created"}).AddRow(
			1, 234, "nice room", "2013-10-02").AddRow(2, 324, "best room", "2014-20-03")

		query := GetRoomsPostgreRequest
		query += " ORDER BY created "
		query += "ASC"

		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(rowsRoom)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewRoomRepository(sqlxDb)
		roomOrder := roomModel.RoomOrder{Sort: "date", Order: "false"}
		rooms, err := rep.GetRooms(roomOrder)
		assert.NoError(t, err)
		assert.Equal(t, roomTest, rooms)
	})
	t.Run("GetRooms3", func(t *testing.T) {
		roomTest := []roomModel.Room{
			{RoomID: 1, Cost: 234, Description: "nice room", Created: "2013-10-02"},
			{RoomID: 2, Cost: 324, Description: "best room", Created: "2014-20-03"},
		}
		rowsRoom := sqlmock.NewRows([]string{"room_id", "cost", "description", "created"}).AddRow(
			1, 234, "nice room", "2013-10-02").AddRow(2, 324, "best room", "2014-20-03")

		query := GetRoomsPostgreRequest
		query += " ORDER BY cost  "
		query += "DESC"

		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(rowsRoom)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewRoomRepository(sqlxDb)
		roomOrder := roomModel.RoomOrder{Sort: "cost", Order: "true"}
		rooms, err := rep.GetRooms(roomOrder)
		assert.NoError(t, err)
		assert.Equal(t, roomTest, rooms)
	})
	t.Run("GetRooms4", func(t *testing.T) {
		roomTest := []roomModel.Room{
			{RoomID: 1, Cost: 234, Description: "nice room", Created: "2013-10-02"},
			{RoomID: 2, Cost: 324, Description: "best room", Created: "2014-20-03"},
		}
		rowsRoom := sqlmock.NewRows([]string{"room_id", "cost", "description", "created"}).AddRow(
			1, 234, "nice room", "2013-10-02").AddRow(2, 324, "best room", "2014-20-03")

		query := GetRoomsPostgreRequest
		query += " ORDER BY cost  "
		query += "ASC"

		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(rowsRoom)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewRoomRepository(sqlxDb)
		roomOrder := roomModel.RoomOrder{Sort: "cost", Order: "false"}
		rooms, err := rep.GetRooms(roomOrder)
		assert.NoError(t, err)
		assert.Equal(t, roomTest, rooms)
	})
	t.Run("GetRoomsError", func(t *testing.T) {

		query := GetRoomsPostgreRequest

		mock.ExpectQuery(query).
			WithArgs().
			WillReturnError(customError.NewCustomError(errors.New("fds"), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewRoomRepository(sqlxDb)
		roomOrder := roomModel.RoomOrder{}
		_, err := rep.GetRooms(roomOrder)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), serverError.ServerInternalError)
	})
}

func TestRoomRepository_CheckRoomExist(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("CheckRoomExist", func(t *testing.T) {
		query := fmt.Sprintf("SELECT exists (%s)", CheckRoomExistPostgreRequest)
		rowsRoom := sqlmock.NewRows([]string{"exists"}).AddRow(
			true)
		roomID := 1
		mock.ExpectQuery(query).
			WithArgs(roomID).
			WillReturnRows(rowsRoom)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewRoomRepository(sqlxDb)

		res, err := rep.CheckRoomExist(roomID)
		assert.NoError(t, err)
		assert.Equal(t, res, true)
	})

	t.Run("CheckRoomExist", func(t *testing.T) {
		query := fmt.Sprintf("SELECT exists (%s)", CheckRoomExistPostgreRequest)

		roomID := 1
		mock.ExpectQuery(query).
			WithArgs(roomID).
			WillReturnError(customError.NewCustomError(errors.New("fds"), serverError.ServerInternalError, 1))

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewRoomRepository(sqlxDb)

		_, err := rep.CheckRoomExist(roomID)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), serverError.ServerInternalError)
	})
}

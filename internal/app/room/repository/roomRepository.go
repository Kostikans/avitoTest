package roomRepository

import (
	"database/sql"
	"fmt"

	roomModel "github.com/Kostikans/avitoTest/internal/app/room/model"
	"github.com/Kostikans/avitoTest/internal/package/serverError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"
	"github.com/jmoiron/sqlx"
)

type RoomRepository struct {
	db *sqlx.DB
}

func NewRoomRepository(db *sqlx.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

func (rRep *RoomRepository) AddRoom(room roomModel.RoomAdd) (roomModel.RoomID, error) {
	roomID := roomModel.RoomID{}
	err := rRep.db.QueryRow(AddRoomPostgreRequest, room.Description, room.Cost).Scan(&roomID.RoomID)
	if err != nil {
		return roomID, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return roomID, nil
}

func (rRep *RoomRepository) DeleteRoom(roomID int) error {
	_, err := rRep.db.Query(DeleteRoomPostgreRequest, roomID)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return err
}

func (rRep *RoomRepository) GenerateQueryForGetRooms(order roomModel.RoomOrder) string {
	query := GetRoomsPostgreRequest
	if order.Sort == "date" {
		query += " ORDER BY created "
		if order.Order == "true" {
			query += "DESC"
		} else if order.Order == "false" {
			query += "ASC"
		}
	} else if order.Sort == "cost" {
		query += " ORDER BY cost  "
		if order.Order == "true" {
			query += "DESC"
		} else if order.Order == "false" {
			query += "ASC"
		}
	}

	return query
}

func (rRep *RoomRepository) GetRooms(order roomModel.RoomOrder) ([]roomModel.Room, error) {
	var rooms []roomModel.Room
	query := rRep.GenerateQueryForGetRooms(order)
	err := rRep.db.Select(&rooms, query)
	if err != nil {
		return rooms, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return rooms, err

}

func (rRep *RoomRepository) CheckRoomExist(roomID int) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT exists (%s)", CheckRoomExistPostgreRequest)
	err := rRep.db.QueryRow(query, roomID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return exists, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return exists, nil
}

package roomRepository

import (
	"errors"

	roomModel "github.com/Kostikans/avitoTest/internal/app/room/model"
	"github.com/Kostikans/avitoTest/internal/package/clientError"
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

func (rRep *RoomRepository) DeleteRoom(roomID int64) error {
	_, err := rRep.db.Exec(DeleteRoomPostgreRequest, roomID)
	if err != nil {
		return customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return err
}

func (rRep *RoomRepository) GenerateQueryForGetRooms(order roomModel.RoomOrder) (string, error) {
	query := GetRoomsPostgreRequest
	if order.DataDesc != "" && order.CostDesc != "" {
		return "", errors.New("available only one order option")
	}

	if order.DataDesc == "true" {
		query += " ORDER BY created DESC"
	} else if order.DataDesc == "false" {
		query += " ORDER BY created ASC"
	}

	if order.CostDesc == "true" {
		query += " ORDER BY cost DESC"
	} else if order.CostDesc == "false" {
		query += " ORDER BY cost ASC"
	}

	return query, nil
}

func (rRep *RoomRepository) GetRooms(order roomModel.RoomOrder) ([]roomModel.Room, error) {
	var rooms []roomModel.Room
	query, err := rRep.GenerateQueryForGetRooms(order)
	if err != nil {
		return rooms, customerror.NewCustomError(err, clientError.BadRequest, 1)
	}
	err = rRep.db.Select(&rooms, query)
	if err != nil {
		return rooms, customerror.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return rooms, err

}

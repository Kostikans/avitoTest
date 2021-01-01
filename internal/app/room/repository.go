package room

import roomModel "github.com/Kostikans/avitoTest/internal/app/room/model"

type Repository interface {
	AddRoom(room roomModel.RoomAdd) (roomModel.RoomID, error)
	DeleteRoom(roomID int64) error
	GetRooms(order roomModel.RoomOrder) ([]roomModel.Room, error)
	CheckRoomExist(roomID int64) (bool, error)
}

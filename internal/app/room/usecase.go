package room

import roomModel "github.com/Kostikans/avitoTest/internal/app/room/model"

type Usecase interface {
	AddRoom(room roomModel.RoomAdd) (roomModel.RoomID, error)
	DeleteRoom(roomID int) error
	GetRooms(order roomModel.RoomOrder) ([]roomModel.Room, error)
	CheckRoomExist(roomID int) (bool, error)
}

package roomUsecase

import (
	"github.com/Kostikans/avitoTest/internal/app/room"
	roomModel "github.com/Kostikans/avitoTest/internal/app/room/model"
)

type RoomUsecase struct {
	RoomRepo room.Repository
}

func NewRoomUsecase(RoomRepo room.Repository) *RoomUsecase {
	return &RoomUsecase{RoomRepo: RoomRepo}
}

func (rUsecase *RoomUsecase) AddRoom(room roomModel.RoomAdd) (roomModel.RoomID, error) {
	return rUsecase.RoomRepo.AddRoom(room)
}

func (rUsecase *RoomUsecase) DeleteRoom(roomID int64) error {
	return rUsecase.RoomRepo.DeleteRoom(roomID)
}

func (rUsecase *RoomUsecase) GetRooms(order roomModel.RoomOrder) ([]roomModel.Room, error) {
	return rUsecase.RoomRepo.GetRooms(order)
}

package roomUsecase

import (
	"errors"

	"github.com/Kostikans/avitoTest/internal/app/room"
	roomModel "github.com/Kostikans/avitoTest/internal/app/room/model"
	"github.com/Kostikans/avitoTest/internal/package/clientError"
	"github.com/Kostikans/avitoTest/internal/package/customError"
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

func (rUsecase *RoomUsecase) DeleteRoom(roomID int) error {
	exist, err := rUsecase.CheckRoomExist(roomID)
	if err != nil {
		return err
	}
	if exist == false {
		return customError.NewCustomError(errors.New("room doesn't exist"), clientError.NotFound, 1)
	}
	return rUsecase.RoomRepo.DeleteRoom(roomID)
}

func (rUsecase *RoomUsecase) GetRooms(order roomModel.RoomOrder) ([]roomModel.Room, error) {
	return rUsecase.RoomRepo.GetRooms(order)
}

func (rUsecase *RoomUsecase) CheckRoomExist(roomID int) (bool, error) {
	return rUsecase.RoomRepo.CheckRoomExist(roomID)
}

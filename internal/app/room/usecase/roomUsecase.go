package roomUsecase

import "github.com/Kostikans/avitoTest/internal/app/room"

type RoomUsecase struct {
	RoomRepo room.Repository
}

func NewRoomUsecase(RoomRepo *room.Repository) *RoomUsecase {
	return &RoomUsecase{RoomRepo: RoomRepo}
}

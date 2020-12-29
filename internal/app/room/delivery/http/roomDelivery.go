package roomDelivery

import (
	"github.com/Kostikans/avitoTest/internal/app/room"
	"github.com/Kostikans/avitoTest/internal/package/logger"
)

type RoomHandler struct {
	roomUsecase room.Usecase
	log         *logger.CustomLogger
}

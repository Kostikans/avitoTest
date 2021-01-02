package roomUsecase

import (
	"errors"
	"testing"

	"github.com/Kostikans/avitoTest/internal/package/serverError"

	"github.com/Kostikans/avitoTest/internal/package/clientError"
	"github.com/Kostikans/avitoTest/internal/package/customError"

	room_mock "github.com/Kostikans/avitoTest/internal/app/room/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRoomUsecase_DeleteRoom(t *testing.T) {
	t.Run("DeleteRoom", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		roomID := 1
		mockRoomRepository := room_mock.NewMockRepository(ctrl)
		mockRoomRepository.EXPECT().
			CheckRoomExist(roomID).
			Return(true, nil)
		mockRoomRepository.EXPECT().
			DeleteRoom(roomID).
			Return(nil)

		roomUsecase := NewRoomUsecase(mockRoomRepository)

		err := roomUsecase.DeleteRoom(roomID)
		assert.NoError(t, err)
	})
	t.Run("DeleteRoomErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		roomID := 1
		mockRoomRepository := room_mock.NewMockRepository(ctrl)
		mockRoomRepository.EXPECT().
			CheckRoomExist(roomID).
			Return(false, nil)

		roomUsecase := NewRoomUsecase(mockRoomRepository)

		err := roomUsecase.DeleteRoom(roomID)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), clientError.NotFound)
	})
	t.Run("DeleteRoomErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		roomID := 1
		mockRoomRepository := room_mock.NewMockRepository(ctrl)
		mockRoomRepository.EXPECT().
			CheckRoomExist(roomID).
			Return(true, customError.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		roomUsecase := NewRoomUsecase(mockRoomRepository)

		err := roomUsecase.DeleteRoom(roomID)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), serverError.ServerInternalError)
	})
}

package roomDelivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Kostikans/avitoTest/internal/package/customError"
	"github.com/Kostikans/avitoTest/internal/package/serverError"

	"github.com/Kostikans/avitoTest/internal/package/clientError"

	"github.com/Kostikans/avitoTest/internal/package/okCodes"

	"github.com/mailru/easyjson"

	room_mock "github.com/Kostikans/avitoTest/internal/app/room/mocks"
	roomModel "github.com/Kostikans/avitoTest/internal/app/room/model"
	"github.com/Kostikans/avitoTest/internal/package/logger"
	"github.com/Kostikans/avitoTest/internal/package/responses"
	"github.com/golang/mock/gomock"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestRoomHandler_AddRoom(t *testing.T) {
	t.Run("AddRoom", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testRoomAdd := roomModel.RoomAdd{Cost: 234, Description: "best room in the world"}
		roomIDTest := roomModel.RoomID{RoomID: 1}
		mockRoomUseCase := room_mock.NewMockUsecase(ctrl)

		mockRoomUseCase.EXPECT().
			AddRoom(testRoomAdd).
			Return(roomIDTest, nil)

		body, err := easyjson.Marshal(testRoomAdd)
		assert.NoError(t, err)
		req, err := http.NewRequest("GET", "rooms/create", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := RoomHandler{
			roomUsecase: mockRoomUseCase,
			log:         logger.NewLogger(os.Stdout),
		}

		handler.AddRoom(rec, req)
		resp := rec.Result()

		roomID := roomModel.RoomID{}

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &roomID)
		assert.NoError(t, err)

		assert.Equal(t, roomID, roomIDTest)
		assert.Equal(t, okCodes.CreateResponse, response.Code)
	})
	t.Run("AddRoomErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testRoomAdd := roomModel.RoomAdd{Cost: 234, Description: "best room in the world"}
		roomIDTest := roomModel.RoomID{RoomID: 1}
		mockRoomUseCase := room_mock.NewMockUsecase(ctrl)

		mockRoomUseCase.EXPECT().
			AddRoom(testRoomAdd).
			Return(roomIDTest, customError.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		body, err := easyjson.Marshal(testRoomAdd)
		assert.NoError(t, err)
		req, err := http.NewRequest("GET", "rooms/create", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := RoomHandler{
			roomUsecase: mockRoomUseCase,
			log:         logger.NewLogger(os.Stdout),
		}

		handler.AddRoom(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)

		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})
	t.Run("AddRoomErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		badRequestTest := "dfsd:fdsxc"
		mockRoomUseCase := room_mock.NewMockUsecase(ctrl)

		body, err := json.Marshal(badRequestTest)
		assert.NoError(t, err)
		req, err := http.NewRequest("GET", "rooms/create", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := RoomHandler{
			roomUsecase: mockRoomUseCase,
			log:         logger.NewLogger(os.Stdout),
		}

		handler.AddRoom(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)

		assert.Equal(t, clientError.BadRequest, response.Code)
	})
}

func TestRoomHandler_GetRooms(t *testing.T) {
	t.Run("GetRoom", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testGetRooms := []roomModel.Room{
			{RoomID: 2, Created: "2013-12-23", Cost: 234, Description: "fds"},
			{RoomID: 3, Created: "2013-12-23", Cost: 234, Description: "fds"},
		}
		mockRoomUseCase := room_mock.NewMockUsecase(ctrl)

		mockRoomUseCase.EXPECT().
			GetRooms(gomock.Any()).
			Return(testGetRooms, nil)

		req, err := http.NewRequest("GET", "rooms/list", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := RoomHandler{
			roomUsecase: mockRoomUseCase,
			log:         logger.NewLogger(os.Stdout),
		}

		handler.GetRooms(rec, req)
		resp := rec.Result()

		var GetRooms []roomModel.Room

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.([]interface{}), &GetRooms)
		assert.NoError(t, err)

		assert.Equal(t, testGetRooms, GetRooms)
		assert.Equal(t, okCodes.OkResponse, response.Code)
	})
	t.Run("GetRoomErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRoomUseCase := room_mock.NewMockUsecase(ctrl)

		mockRoomUseCase.EXPECT().
			GetRooms(gomock.Any()).
			Return(nil, customError.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		req, err := http.NewRequest("GET", "rooms/list", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := RoomHandler{
			roomUsecase: mockRoomUseCase,
			log:         logger.NewLogger(os.Stdout),
		}

		handler.GetRooms(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})
}

func TestRoomHandler_DeleteRoom(t *testing.T) {
	t.Run("DeleteRoom", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRoomUseCase := room_mock.NewMockUsecase(ctrl)
		roomID := 2
		mockRoomUseCase.EXPECT().
			DeleteRoom(roomID).
			Return(nil)

		req, err := http.NewRequest("GET", "rooms/delete?room_id=2", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := RoomHandler{
			roomUsecase: mockRoomUseCase,
			log:         logger.NewLogger(os.Stdout),
		}

		handler.DeleteRoom(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)

		assert.Equal(t, okCodes.OkResponse, response.Code)
	})
	t.Run("DeleteRoom", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRoomUseCase := room_mock.NewMockUsecase(ctrl)
		roomID := 2
		mockRoomUseCase.EXPECT().
			DeleteRoom(roomID).
			Return(customError.NewCustomError(errors.New(""), serverError.ServerInternalError, 1))

		req, err := http.NewRequest("GET", "rooms/delete?room_id=2", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := RoomHandler{
			roomUsecase: mockRoomUseCase,
			log:         logger.NewLogger(os.Stdout),
		}

		handler.DeleteRoom(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)

		assert.Equal(t, serverError.ServerInternalError, response.Code)
	})
}

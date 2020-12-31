package roomDelivery

import (
	"net/http"
	"strconv"

	"github.com/Kostikans/avitoTest/internal/package/responses"

	"github.com/Kostikans/avitoTest/internal/package/clientError"

	"github.com/Kostikans/avitoTest/internal/package/customError"

	roomModel "github.com/Kostikans/avitoTest/internal/app/room/model"
	"github.com/mailru/easyjson"

	"github.com/Kostikans/avitoTest/internal/app/room"
	"github.com/Kostikans/avitoTest/internal/package/logger"
	"github.com/gorilla/mux"
)

type RoomHandler struct {
	roomUsecase room.Usecase
	log         *logger.CustomLogger
}

func NewRoomHandler(r *mux.Router, usecase room.Usecase, log *logger.CustomLogger) *RoomHandler {
	handler := RoomHandler{roomUsecase: usecase, log: log}

	r.HandleFunc("/api/v1/rooms", handler.AddRoom).Methods("POST")
	r.HandleFunc("/api/v1/rooms", handler.GetRooms).Methods("GET")
	r.HandleFunc("/api/v1/rooms/{room_id:[0-9]+}", handler.DeleteRoom).Methods("DELETE")
	return &handler
}

func (rh *RoomHandler) AddRoom(w http.ResponseWriter, r *http.Request) {
	roomAdd := roomModel.RoomAdd{}
	err := easyjson.UnmarshalFromReader(r.Body, &roomAdd)
	if err != nil {
		customError.PostError(w, r, rh.log, err, clientError.BadRequest)
		return
	}

	roomID, err := rh.roomUsecase.AddRoom(roomAdd)
	if err != nil {
		customError.PostError(w, r, rh.log, err, nil)
		return
	}
	responses.SendDataResponse(w, roomID)
}

func (rh *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	roomIDVar := mux.Vars(r)["room_id"]
	roomID, err := strconv.ParseInt(roomIDVar, 10, 64)
	if err != nil {
		customError.PostError(w, r, rh.log, err, clientError.BadRequest)
		return
	}

	err = rh.roomUsecase.DeleteRoom(roomID)
	if err != nil {
		customError.PostError(w, r, rh.log, err, nil)
		return
	}
	responses.SendOkResponse(w)
}

func (rh *RoomHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
	dateDesc := r.FormValue("dateDesc")
	costDesc := r.FormValue("costDesc")

	rooms, err := rh.roomUsecase.GetRooms(roomModel.RoomOrder{CostDesc: costDesc, DataDesc: dateDesc})
	if err != nil {
		customError.PostError(w, r, rh.log, err, nil)
		return
	}
	responses.SendDataResponse(w, rooms)
}

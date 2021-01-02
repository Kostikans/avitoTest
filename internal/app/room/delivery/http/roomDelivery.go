package roomDelivery

import (
	"net/http"
	"strconv"

	"github.com/Kostikans/avitoTest/internal/package/okCodes"

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

	r.HandleFunc("/rooms/create", handler.AddRoom).Methods("POST")
	r.HandleFunc("/rooms/list", handler.GetRooms).Methods("GET")
	r.Path("/rooms/delete").Queries("room_id", "{room_id:[0-9]+}").HandlerFunc(handler.DeleteRoom).Methods("DELETE")
	return &handler
}

// swagger:route POST /rooms/create rooms AddRoom
// responses:
//  201: roomID
//  400: badrequest
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
	responses.SendDataResponse(w, okCodes.CreateResponse, roomID)
}

// swagger:route DELETE /rooms/delete rooms DeleteRoom
// responses:
//  400: badrequest
//  404: notfound
func (rh *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	roomIDVar := r.FormValue("room_id")
	roomID, err := strconv.Atoi(roomIDVar)
	if err != nil {
		customError.PostError(w, r, rh.log, err, clientError.BadRequest)
		return
	}

	err = rh.roomUsecase.DeleteRoom(roomID)
	if err != nil {
		customError.PostError(w, r, rh.log, err, nil)
		return
	}
	responses.SendOkResponse(w, okCodes.OkResponse)
}

// swagger:route GET /rooms/list rooms GetRooms
// responses:
// 200: rooms
func (rh *RoomHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
	sort := r.FormValue("sort")
	order := r.FormValue("desc")

	rooms, err := rh.roomUsecase.GetRooms(roomModel.RoomOrder{Sort: sort, Order: order})
	if err != nil {
		customError.PostError(w, r, rh.log, err, nil)
		return
	}
	responses.SendDataResponse(w, okCodes.OkResponse, rooms)
}

package roomDelivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/schema"

	"github.com/Kostikans/avitoTest/internal/package/okCodes"

	"github.com/Kostikans/avitoTest/internal/package/responses"

	"github.com/Kostikans/avitoTest/internal/package/clientError"

	"github.com/Kostikans/avitoTest/internal/package/customError"

	"github.com/Kostikans/avitoTest/internal/app/room"
	roomModel "github.com/Kostikans/avitoTest/internal/app/room/model"
	"github.com/Kostikans/avitoTest/internal/package/logger"
	"github.com/gorilla/mux"
)

type RoomHandler struct {
	roomUsecase room.Usecase
	log         *logger.CustomLogger
	decoder     *schema.Decoder
}

func NewRoomHandler(r *mux.Router, usecase room.Usecase, log *logger.CustomLogger, decoder *schema.Decoder) *RoomHandler {
	handler := RoomHandler{roomUsecase: usecase, log: log, decoder: decoder}

	r.HandleFunc("/rooms/create", handler.AddRoom).Methods("POST")
	r.HandleFunc("/rooms/list", handler.GetRooms).Methods("GET")
	r.Path("/rooms/delete").Queries("room_id", "{room_id}").HandlerFunc(handler.DeleteRoom).Methods("DELETE")
	return &handler
}

// swagger:route POST /rooms/create rooms AddRoom
// responses:
//  201: roomID
//  400: badrequest
func (rh *RoomHandler) AddRoom(w http.ResponseWriter, r *http.Request) {
	roomAdd := roomModel.RoomAdd{}
	err := r.ParseForm()
	if err != nil {
		customError.PostError(w, r, rh.log, errors.New("cannot parse form"), clientError.BadRequest)
		return
	}
	err = rh.decoder.Decode(&roomAdd, r.PostForm)
	if err != nil {
		customError.PostError(w, r, rh.log, errors.New("cannot decode into struct"), clientError.BadRequest)
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

package swagger

type RoomAddWrap struct {
	Cost        int    `json:"cost"`
	Description string `json:"description"`
}

type RoomIDWrap struct {
	RoomID int `json:"room_id"`
}

type RoomWrap struct {
	RoomID      int    `json:"room_id" `
	Created     string `json:"create" `
	Cost        int    `json:"cost" `
	Description string `json:"description"`
}

//swagger:parameters AddRoom
type AddRoomRequestWrap struct {
	//in: body
	RoomAdd RoomAddWrap
}

//swagger:parameters GetRooms
type GetRoomsRequestWrap struct {
	//sort param; "date","cost"
	//in: query
	Sort string `json:"sort"`
	//order param; "true" - desc, "false" - asc
	//in: query
	Desc bool `json:"desc"`
}

//swagger:response rooms
type GetRoomsResponseWrap struct {
	//in: body
	RoomAdd []RoomWrap
}

//swagger:parameters DeleteRoom
type DeleteRoomRequestWrap struct {
	//in: query
	//required: true
	RoomID int `json:"room_id"`
}

//swagger:response roomID
type AddRoomResponseWrap struct {
	//in: body
	RoomID RoomIDWrap
}

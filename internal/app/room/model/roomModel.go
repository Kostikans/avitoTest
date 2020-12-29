package roomModel

import "time"

type Room struct {
	RoomID      int64     `json:"room_id"`
	Created     time.Time `json:"create"`
	Cost        int64     `json:"cost"`
	Description string    `json:"description"`
}

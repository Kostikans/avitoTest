//go:generate easyjson -all roomModel.go
package roomModel

// easyjson:json
type Room struct {
	RoomID      int    `json:"room_id" db:"room_id"  mapstructure:"room_id"`
	Created     string `json:"create" db:"created"  mapstructure:"create"`
	Cost        int    `json:"cost" db:"cost"  mapstructure:"cost"`
	Description string `json:"description" db:"description"  mapstructure:"description"`
}

// easyjson:json
type RoomAdd struct {
	Cost        int    `json:"cost" schema:"cost,required"`
	Description string `json:"description" schema:"description,required"`
}

type RoomOrder struct {
	Sort  string
	Order string
}

// easyjson:json
type RoomID struct {
	RoomID int `json:"room_id" mapstructure:"room_id"`
}

//go:generate easyjson -all bookingModel.go
package bookingModel

// easyjson:json
type Booking struct {
	BookingID int    `json:"booking_id" db:"booking_id" mapstructure:"booking_id"`
	DateStart string `json:"date_start" db:"date_start" mapstructure:"date_start"`
	DateEnd   string `json:"date_end" db:"date_end" mapstructure:"date_end"`
}

// easyjson:json
type BookingAdd struct {
	RoomID    int    `json:"room_id"`
	DateStart string `json:"date_start" db:"date_start" `
	DateEnd   string `json:"date_end" db:"date_end" `
}

// easyjson:json
type BookingID struct {
	BookingID int `json:"booking_id" mapstructure:"booking_id"`
}

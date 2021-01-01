//go:generate easyjson -all bookingModel.go
package bookingModel

// easyjson:json
type Booking struct {
	BookingID int64  `json:"booking_id" db:"booking_id" `
	DateStart string `json:"date_start" db:"date_start"`
	DateEnd   string `json:"date_end" db:"date_end" `
}

// easyjson:json
type BookingAdd struct {
	RoomID    int64  `json:"room_id"`
	DateStart string `json:"date_start" db:"date_start" validate:"required,datetime"`
	DateEnd   string `json:"date_end" db:"date_end" validate:"required,datetime"`
}

// easyjson:json
type BookingID struct {
	BookingID int64 `json:"booking_id"`
}

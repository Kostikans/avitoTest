package swagger

type BookingWrap struct {
	BookingID int    `json:"booking_id"`
	DateStart string `json:"date_start"`
	DateEnd   string `json:"date_end"`
}

type BookingIDWrap struct {
	BookingID int `json:"booking_id"`
}

//swagger:parameters AddBooking
type AddBookingRequestWrap struct {
	///in: formData
	//required: true
	RoomID int `json:"room_id"`
	//in: formData
	//required: true
	DateStart string `json:"date_start"`
	//in: formData
	//required: true
	DateEnd string `json:"date_end"`
}

//swagger:response bookings
type GetBookingsResponseWrap struct {
	//in: body
	Bookings []BookingWrap
}

//swagger:parameters GetBookings
type GetBookingsRequestWrap struct {
	//in: query
	//required: true
	BookingID int `json:"room_id"`
}

//swagger:parameters DeleteBooking
type DeleteBookingRequestWrap struct {
	//in: query
	//required: true
	BookingID int `json:"booking_id"`
}

//swagger:response bookingID
type AddBookingResponseWrap struct {
	//in: body
	BookingAdd BookingIDWrap
}

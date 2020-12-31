package bookingRepository

const AddBookingPostgreRequest = "INSERT INTO booking(room_id,date_start,date_end) VALUES($1,$2,$3) RETURNING booking_id"

const DeleteRoomPostgreRequest = "DELETE FROM booking WHERE booking_id=$1"

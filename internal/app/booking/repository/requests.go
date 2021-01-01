package bookingRepository

const AddBookingPostgreRequest = "INSERT INTO booking(room_id,date_start,date_end) VALUES($1,$2,$3) RETURNING booking_id"

const DeleteRoomPostgreRequest = "DELETE FROM booking WHERE booking_id=$1"

const GetBookingsPostgreRequest = "SELECT booking_id,to_char(date_start,'YYYY-MM-DD') as date_start ,to_char(date_end,'YYYY-MM-DD') as date_end" +
	" FROM booking WHERE room_id=$1 ORDER BY date_start"

const CheckBookingExistPostgreRequest = "SELECT * from booking WHERE booking_id=$1"

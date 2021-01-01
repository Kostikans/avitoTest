package roomRepository

const AddRoomPostgreRequest = "INSERT INTO rooms(description,cost) VALUES($1,$2) RETURNING room_id"

const DeleteRoomPostgreRequest = "DELETE FROM rooms WHERE room_id=$1"

const GetRoomsPostgreRequest = "SELECT room_id,description,cost,to_char(created,'YYYY-MM-DD') as created FROM rooms"

const CheckRoomExistPostgreRequest = "SELECT * from rooms WHERE room_id=$1"

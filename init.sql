create table rooms
(
    room_id     serial not null primary key,
    description text,
    cost        int,
    created     timestamp default now()
);

create table booking
(
    booking_id  serial not null primary key,
    room_id     int,
    date_start  timestamp,
    date_end    timestamp,
    CONSTRAINT fk_rooms FOREIGN KEY(room_id) REFERENCES rooms(room_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE INDEX booking_idx_1 ON booking(room_id);
CREATE INDEX booking_where__idx ON booking(room_id,booking_id,date_start,date_end); --index only scan


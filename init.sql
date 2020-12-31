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
    CONSTRAINT fk_rooms
        FOREIGN KEY(room_id)
            REFERENCES rooms(room_id)
            ON DELETE CASCADE
);
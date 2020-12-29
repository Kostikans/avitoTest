package roomRepository

import (
	"github.com/jmoiron/sqlx"
)

type RoomRepository struct {
	db *sqlx.DB
}

func NewRoomRepository(db *sqlx.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

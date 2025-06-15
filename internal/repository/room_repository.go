package repository

import (
	"booking-app/internal/entity"
	"database/sql"
	"log"
)

type RoomRepository struct {
	DB *sql.DB
}

func NewRoomRepository(db *sql.DB) *RoomRepository {
	return &RoomRepository{
		DB: db,
	}
}

type RoomRepositoryInterface interface {
	GetAllActiveRoom() ([]entity.Room, error)
}

func (s *RoomRepository) GetAllActiveRoom() ([]entity.Room, error) {
	result, err := s.DB.Query("select room_id, room_name, room_description, room_status from room where room_status = 'Active'")

	if err != nil {
		return nil, err
	}

	defer result.Close()

	var rooms []entity.Room

	for result.Next() {
		var room entity.Room
		err := result.Scan(&room.RoomId, &room.Name, &room.Description, &room.Status)
		if err != nil {
			log.Println("Error scanning room: ", err.Error())
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

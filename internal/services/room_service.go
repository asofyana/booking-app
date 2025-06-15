package services

import (
	"booking-app/internal/entity"
	"booking-app/internal/repository"

	"github.com/gin-gonic/gin"
)

type RoomService struct {
	RoomRepository repository.RoomRepositoryInterface
}

func NewRoomService(roomRepository repository.RoomRepositoryInterface) *RoomService {
	return &RoomService{
		RoomRepository: roomRepository,
	}
}

type RoomServiceInterface interface {
	GetAllActiveRoom(c *gin.Context) ([]entity.Room, error)
}

func (s *RoomService) GetAllActiveRoom(c *gin.Context) ([]entity.Room, error) {
	return s.RoomRepository.GetAllActiveRoom()
}

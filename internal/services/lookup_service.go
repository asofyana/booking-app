package services

import (
	"booking-app/internal/entity"
	"booking-app/internal/repository"

	"github.com/gin-gonic/gin"
)

type LookupService struct {
	LookupRepository repository.LookupRepositoryInterface
}

func NewLookupService(lookupRepository repository.LookupRepositoryInterface) *LookupService {
	return &LookupService{
		LookupRepository: lookupRepository,
	}
}

type LookupServiceInterface interface {
	GetAllActiveLookupByType(c *gin.Context, ltype string) ([]entity.Lookup, error)
}

func (s *LookupService) GetAllActiveLookupByType(c *gin.Context, ltype string) ([]entity.Lookup, error) {
	return s.LookupRepository.GetAllActiveLookupByType(ltype)
}

package repository

import (
	"booking-app/internal/entity"
	"database/sql"
	"log"
)

type LookupRepository struct {
	DB *sql.DB
}

func NewLookupRepository(db *sql.DB) *LookupRepository {
	return &LookupRepository{
		DB: db,
	}
}

type LookupRepositoryInterface interface {
	GetAllActiveLookupByType(ltype string) ([]entity.Lookup, error)
}

func (s *LookupRepository) GetAllActiveLookupByType(ltype string) ([]entity.Lookup, error) {

	result, err := s.DB.Query("select lookup_type, lookup_code, lookup_text, lookup_status from lookup where lookup_status = 'Active' and lookup_type=?", ltype)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	var lookups []entity.Lookup

	for result.Next() {
		var lookup entity.Lookup
		err := result.Scan(&lookup.LookupType, &lookup.LookupCode, &lookup.LookupText, &lookup.LookupStatus)
		if err != nil {
			log.Println("Error scanning room: ", err.Error())
			return nil, err
		}
		lookups = append(lookups, lookup)
	}

	return lookups, nil

}

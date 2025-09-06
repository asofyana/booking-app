package entity

import "time"

type User struct {
	UserId      int       `json:"user_id"`
	UserName    string    `json:"username"`
	Password    string    `json:"password"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	CreatedDate time.Time `json:"created_date"`
	CreatedBy   string    `json:"created_by"`
	UpdatedDate time.Time `json:"updated_date"`
	UpdatedBy   string    `json:"updated_by"`
}

type Room struct {
	RoomId      int       `json:"room_id"`
	Name        string    `json:"room_name"`
	Description string    `json:"room_description"`
	Status      string    `json:"room_status"`
	StatusText  string    `json:"status_text"`
	CreatedDate time.Time `json:"created_date"`
	CreatedBy   string    `json:"created_by"`
	UpdatedDate time.Time `json:"updated_date"`
	UpdatedBy   string    `json:"updated_by"`
}

type Booking struct {
	BookingId        int       `json:"booking_id"`
	Title            string    `json:"title"`
	RoomId           int       `json:"room_id"`
	RoomName         string    `json:"room_name"`
	Description      string    `json:"booking_description"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	ParticipantCount int       `json:"participant_count"`
	Status           string    `json:"status"`
	StatusText       string    `json:"status_text"`
	Activity         string    `json:"activity_code"`
	ActivityText     string    `json:"activity_text"`
	CreatedDate      time.Time `json:"created_date"`
	CreatedBy        string    `json:"created_by"`
	UpdatedDate      time.Time `json:"updated_date"`
	UpdatedBy        string    `json:"updated_by"`
}

type Lookup struct {
	LookupType   string `json:"lookup_type"`
	LookupCode   string `json:"lookup_code"`
	LookupText   string `json:"lookup_text"`
	LookupStatus string `json:"lookup_status"`
}

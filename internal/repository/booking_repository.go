package repository

import (
	"booking-app/internal/entity"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"booking-app/internal/utils"
)

type BookingRepository struct {
	DB *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{
		DB: db,
	}
}

type BookingRepositoryInterface interface {
	GetAllIncomingBooking() ([]entity.Booking, error)
	InsertBooking(booking entity.Booking) error
	GetOverlapBookingCount(booking entity.Booking) int
	GetIncomingBookingByUsername(username string) ([]entity.Booking, error)
	GetBookingById(bookingId int) (entity.Booking, error)
	UpdateBookingStatus(bookingId int, status string) error
	SearchBooking(booking entity.Booking) ([]entity.Booking, error)
	GetAllIncomingBookingDashboard() ([]entity.Booking, error)
	RejectBooking(bookingId int, rejectReason string) error
}

func (s *BookingRepository) GetAllIncomingBooking() ([]entity.Booking, error) {
	result, err := s.DB.Query("select a.booking_id, a.title, a.start_date, a.end_date, a.participant_count, a.status, a.activity_code, c.lookup_text as activity_text, a.organizer, a.pic, a.pic_contactno, IFNULL(a.reject_reason,'') as reject_reason, a.created_by " +
		"from booking a, lookup c where a.activity_code = c.lookup_code and c.lookup_type='ACTIVITY' and a.start_date > current_timestamp order by a.start_date")

	if err != nil {
		return nil, err
	}
	defer result.Close()

	var bookings []entity.Booking

	for result.Next() {
		var booking entity.Booking
		err := result.Scan(&booking.BookingId, &booking.Title, &booking.StartDate, &booking.EndDate, &booking.ParticipantCount, &booking.Status, &booking.Activity, &booking.ActivityText, &booking.Organizer, &booking.Pic, &booking.PicContactNo, &booking.RejectReason, &booking.CreatedBy)
		if err != nil {
			log.Println("Error scanning booking: ", err.Error())
			return nil, err
		}
		rooms, _ := getRoomByBookingId(s.DB, booking.BookingId)
		booking.Rooms = rooms
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (s *BookingRepository) InsertBooking(booking entity.Booking) error {
	result, err := s.DB.Exec("INSERT INTO booking (title,start_date,end_date,participant_count,status,activity_code,organizer,pic,pic_contactno,created_by) VALUES (?,?,?,?,?,?,?,?,?,?)",
		booking.Title, booking.StartDate, booking.EndDate, booking.ParticipantCount, booking.Status, booking.Activity, booking.Organizer, booking.Pic, booking.PicContactNo, booking.CreatedBy)

	if err != nil {
		log.Println("Error inserting booking: ", err.Error())
		return err
	}

	// Insert booking room
	bookingId, _ := result.LastInsertId()
	for _, room := range booking.Rooms {
		_, err := s.DB.Exec("insert into booking_room(booking_id, room_id) values(?,?)", bookingId, room.RoomId)
		if err != nil {
			log.Println("Error inserting booking_room: ", err.Error())
			return err
		}
	}

	return nil
}

func (s *BookingRepository) GetOverlapBookingCount(booking entity.Booking) int {
	total := 0
	for _, room := range booking.Rooms {
		var count int
		err := s.DB.QueryRow("select count(1) from booking a, booking_room b where a.booking_id = b.booking_id and b.room_id = ? and start_date < ? and end_date > ?", room.RoomId, booking.EndDate, booking.StartDate).Scan(&count)

		if err != nil {
			fmt.Println(err)
			return -1
		}
		total += count
	}
	return total
}

func (s *BookingRepository) GetIncomingBookingByUsername(username string) ([]entity.Booking, error) {
	result, err := s.DB.Query("select a.booking_id, a.title, a.start_date, a.end_date, a.participant_count, a.status, a.activity_code, c.lookup_text as activity_text, a.organizer, a.pic, a.pic_contactno, IFNULL(a.reject_reason,'') as reject_reason "+
		"from booking a, lookup c where a.activity_code = c.lookup_code and c.lookup_type='ACTIVITY' "+
		"and a.created_by = ? and a.start_date > current_timestamp order by a.start_date", username)

	if err != nil {
		return nil, err
	}
	defer result.Close()

	var bookings []entity.Booking

	for result.Next() {
		var booking entity.Booking
		err := result.Scan(&booking.BookingId, &booking.Title, &booking.StartDate, &booking.EndDate, &booking.ParticipantCount, &booking.Status, &booking.Activity, &booking.ActivityText, &booking.Organizer, &booking.Pic, &booking.PicContactNo, &booking.RejectReason)
		if err != nil {
			log.Println("Error scanning booking: ", err.Error())
			return nil, err
		}
		rooms, _ := getRoomByBookingId(s.DB, booking.BookingId)
		booking.Rooms = rooms
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (s *BookingRepository) GetBookingById(bookingId int) (entity.Booking, error) {
	result, err := s.DB.Query("select a.booking_id, a.title, a.start_date, a.end_date, a.participant_count, a.status, a.activity_code, c.lookup_text as activity_text, a.organizer, a.pic, a.pic_contactno, IFNULL(a.reject_reason,'') as reject_reason, a.created_by "+
		"from booking a, lookup c where a.booking_id = ? and a.activity_code = c.lookup_code and c.lookup_type='ACTIVITY'", bookingId)

	if err != nil {
		log.Println("Error query booking: ", err.Error())
		return entity.Booking{}, err
	}
	defer result.Close()

	var booking entity.Booking

	if result.Next() {
		err2 := result.Scan(&booking.BookingId, &booking.Title, &booking.StartDate, &booking.EndDate, &booking.ParticipantCount, &booking.Status, &booking.Activity, &booking.ActivityText, &booking.Organizer, &booking.Pic, &booking.PicContactNo, &booking.RejectReason, &booking.CreatedBy)
		if err2 != nil {
			log.Println("Error scanning booking: ", err2.Error())
			return entity.Booking{}, err2
		}
		rooms, _ := getRoomByBookingId(s.DB, booking.BookingId)
		booking.Rooms = rooms
		return booking, nil
	}
	return entity.Booking{}, errors.New("invalid bookingId")
}

func (s *BookingRepository) UpdateBookingStatus(bookingId int, status string) error {
	_, err := s.DB.Exec("UPDATE booking set status = ? where booking_id = ?",
		status, bookingId)

	if err != nil {
		log.Println("Error updating booking status: ", err.Error())
		return err
	}

	return nil

}

func getRoomByBookingId(DB *sql.DB, bookingId int) ([]entity.Room, error) {
	result, err := DB.Query("select b.room_id, b.room_name, b.room_description from booking_room a, room b where a.booking_id = ? and a.room_id = b.room_id", bookingId)

	if err != nil {
		return nil, err
	}
	defer result.Close()

	var rooms []entity.Room

	for result.Next() {
		var room entity.Room
		err := result.Scan(&room.RoomId, &room.Name, &room.Description)
		if err != nil {
			log.Println("Error scanning room: ", err.Error())
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (s *BookingRepository) SearchBooking(booking entity.Booking) ([]entity.Booking, error) {
	sql := "select a.booking_id, a.title, a.start_date, a.end_date, a.participant_count, a.status, a.activity_code, c.lookup_text as activity_text, a.organizer, a.pic, a.pic_contactno, IFNULL(a.reject_reason,'') as reject_reason " +
		"from booking a, lookup c where a.activity_code = c.lookup_code and c.lookup_type='ACTIVITY'"
	conditions := []string{}
	params := []interface{}{}

	if !booking.StartDate.IsZero() {
		conditions = append(conditions, "a.start_date >= ?")
		params = append(params, booking.StartDate.Format("2006-01-02")+" 00:00:00")
	}

	if !booking.EndDate.IsZero() {
		conditions = append(conditions, "a.start_date < ?")
		params = append(params, booking.EndDate.AddDate(0, 0, 1).Format("2006-01-02")+" 00:00:00")
	}

	if len(conditions) > 0 {
		sql += " AND " + strings.Join(conditions, " AND ")
	}

	result, err := s.DB.Query(sql, params...)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer result.Close()

	var bookings []entity.Booking

	for result.Next() {
		var booking entity.Booking
		err := result.Scan(&booking.BookingId, &booking.Title, &booking.StartDate, &booking.EndDate, &booking.ParticipantCount, &booking.Status, &booking.Activity, &booking.ActivityText, &booking.Organizer, &booking.Pic, &booking.PicContactNo, &booking.RejectReason)
		if err != nil {
			log.Println("Error scanning booking: ", err.Error())
			return nil, err
		}
		rooms, _ := getRoomByBookingId(s.DB, booking.BookingId)
		booking.Rooms = rooms
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (s *BookingRepository) GetAllIncomingBookingDashboard() ([]entity.Booking, error) {
	sql := "select a.title, a.start_date, a.end_date, c.room_name, c.css_class " +
		"from booking a inner join booking_room b on a.booking_id = b.booking_id " +
		"inner join room c on b.room_id = c.room_id where a.start_date > datetime('now', ?) " +
		"and a.status in ('Approved', 'Pending') order by a.start_date"

	result, err := s.DB.Query(sql, utils.GetConfig().PrevDays)

	if err != nil {
		return nil, err
	}
	defer result.Close()

	var bookings []entity.Booking

	for result.Next() {
		var booking entity.Booking
		err := result.Scan(&booking.Title, &booking.StartDate, &booking.EndDate, &booking.RoomName, &booking.CssClass)
		if err != nil {
			log.Println("Error scanning booking: ", err.Error())
			return nil, err
		}
		rooms, _ := getRoomByBookingId(s.DB, booking.BookingId)
		booking.Rooms = rooms
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (s *BookingRepository) RejectBooking(bookingId int, rejectReason string) error {
	_, err := s.DB.Exec("UPDATE booking set status = 'Rejected', reject_reason = ? where booking_id = ?",
		rejectReason, bookingId)

	if err != nil {
		log.Println("Error reject booking: ", err.Error())
		return err
	}

	return nil

}

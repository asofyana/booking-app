# Room Booking Application

A web-based room booking management system built with Go and SQLite.

## Features

- **User Authentication**: Session-based authentication with bcrypt password hashing
- **Room Management**: Multi-room booking support with overlap detection
- **Booking Workflow**: Create, view, approve/reject bookings
- **Admin Portal**: User management, booking download
- **Dashboard**: Visual calendar view of upcoming bookings
- **Multi-language**: English and Indonesian support

## Tech Stack

| Component | Technology |
|-----------|------------|
| Web Framework | Gin |
| Database | SQLite |
| Authentication | Session-based (gorilla/sessions) |
| Password Hashing | bcrypt |
| Frontend | Bootstrap 5, jQuery, FullCalendar |

## Installation

1. Clone the repository
2. Ensure Go 1.23+ is installed
3. Build and run:

```bash
go build -o booking-app main.go
./booking-app
```

Or run directly:

```bash
go run main.go
```

The application will start on `http://localhost:8080`

## Configuration

Edit `config.yaml` to customize:

```yaml
port: "8080"                    # Server port
defaultPassword: "$2a$14$..."   # Default user password (bcrypt hash)
defaultLanguage: "id-ID"        # Default language (en-US or id-ID)
dbFile: "booking.db"            # SQLite database file
prevDays: "-30 days"            # Dashboard lookback period
```

## Database Schema

The database is auto-created on first run. The schema is in `scripts/schema.sql`:

- `USERS` - User accounts (username, password, role, status)
- `ROOM` - Room inventory
- `BOOKING` - Booking records
- `BOOKING_ROOM` - Many-to-many mapping
- `LOOKUP` - Reference data (activities, roles, statuses)

## Usage

### Default Login

Default user credentials:
- **Username**: `admin` or `user`
- **Password**: Use the hash from `config.yaml` (default: `admin123`)

### Routes

| Route | Description |
|-------|-------------|
| `/` | Login page |
| `/home` | Dashboard with calendar view |
| `/booking/home` | Booking home (my bookings + all bookings) |
| `/booking/create` | Create new booking |
| `/booking/view?id=X` | View booking details |
| `/booking/approval?id=X` | Approve/reject booking (admin only) |
| `/admin/user-search` | User management (admin only) |
| `/user/change-password` | Change password |

## Project Structure

```
internal/
├── app/          # App initialization and DI
├── db/           # Database connection
├── entity/       # Data models
├── handler/      # HTTP handlers
├── repository/   # Data access layer
├── services/     # Business logic
└── utils/        # Utilities (auth, translation, config)
```

## Development

### Building

```bash
go build -o booking-app main.go
```

### Database Access

```bash
sqlite3 booking.db
.tables
```

## Security Notes

- Session cookies are `HttpOnly` with 15-minute expiry
- Passwords are hashed with bcrypt (cost 14)
- Change the session secret in `auth_service.go` before production use
- Set `Secure: true` for cookies in HTTPS environments

## License

Internal use only.

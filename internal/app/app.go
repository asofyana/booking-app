package app

import (
	"booking-app/internal/db"
	"booking-app/internal/handler"
	"booking-app/internal/repository"
	"booking-app/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() error {
	dbConn := db.NewSqliteConnection("./booking.db")

	err := dbConn.Connect()
	if err != nil {
		log.Fatal("Error connecting to database ", err.Error())
		return err
	}
	defer dbConn.Close()

	userRepository := repository.NewUserRepository(dbConn.DB)
	loginService := services.NewLoginService(userRepository)

	loginHandler := handler.NewLoginHandler(loginService)

	roomRepository := repository.NewRoomRepository(dbConn.DB)
	bookingRepository := repository.NewBookingRepository(dbConn.DB)
	roomService := services.NewRoomService(roomRepository)
	bookingService := services.NewBookingService(bookingRepository)
	bookingHandler := handler.NewBookingHandler(roomService, bookingService)

	app := gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	app.LoadHTMLGlob("templates/*")

	app.Static("/assets", "./static")

	gin.SetMode(gin.DebugMode)
	handler.InitializeRoutes(app)
	handler.RegisterRouteLogin(app, loginHandler)
	handler.RegisterRouteBooking(app, *bookingHandler)

	return app.Run(":8080")
}

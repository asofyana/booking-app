package main

import (
	"booking-app/internal/app"
	"log"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	if err := app.NewApp().Start(); err != nil {
		log.Fatal("failed start: ", err.Error())
	}

}

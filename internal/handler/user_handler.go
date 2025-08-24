package handler

import (
	"booking-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService services.UserServiceInterface
}

func NewUserHandler(userService services.UserServiceInterface) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (s *UserHandler) ShowChangePasswordPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"change-password.html",
		gin.H{
			"title": "Change Password",
		},
	)
}

func (s *UserHandler) ProcessChangePassword(c *gin.Context) {
	oldPassword := c.PostForm("old-password")
	newPassword := c.PostForm("new-password")
	confirmPassword := c.PostForm("confirm-password")

	err := s.UserService.UpdatePassword(c, oldPassword, newPassword, confirmPassword)

	message := "Success"
	alert := "alert-success"
	if err != nil {
		alert = "alert-danger"
		message = err.Error()
	}

	c.HTML(http.StatusOK, "change-password.html", gin.H{
		"title": "Change Password", "message": message, "alert": alert})

}

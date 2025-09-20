package handler

import (
	"booking-app/internal/services"
	"booking-app/internal/utils"
	"fmt"
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
	user := services.GetUserSession(c)
	c.HTML(
		http.StatusOK,
		"change-password.html",
		gin.H{
			"title": "Change Password", "User": user,
		},
	)
}

func (s *UserHandler) ProcessChangePassword(c *gin.Context) {
	user := services.GetUserSession(c)

	oldPassword := c.PostForm("old-password")
	newPassword := c.PostForm("new-password")
	confirmPassword := c.PostForm("confirm-password")

	err := s.UserService.UpdatePassword(c, oldPassword, newPassword, confirmPassword)

	message := utils.Translate("change_password_success", nil)
	alert := "alert-success"
	if err != nil {
		fmt.Println(err)
		alert = "alert-danger"
		message = utils.Translate("change_password_failed", nil)
	}

	c.HTML(http.StatusOK, "change-password.html", gin.H{
		"title": "Change Password", "User": user, "message": message, "alert": alert})

}

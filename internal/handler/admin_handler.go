package handler

import (
	"booking-app/internal/entity"
	"booking-app/internal/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	userService services.UserServiceInterface
}

func NewAdminHandler(userService services.UserServiceInterface) *AdminHandler {
	return &AdminHandler{
		userService: userService,
	}
}

func (h *AdminHandler) UserSearch(c *gin.Context) {
	user := services.GetUserSession(c)

	c.HTML(http.StatusOK, "user-search.html", gin.H{
		"title": "User Search", "User": user,
	})
}

func (h *AdminHandler) UserSearchPost(c *gin.Context) {
	var userSearch entity.User

	user := services.GetUserSession(c)

	name := c.Request.FormValue("fullName")
	username := c.Request.FormValue("username")

	fmt.Println("name:", name)
	fmt.Println("username:", username)

	userSearch.FullName = name
	userSearch.UserName = username

	userSearchList, _ := h.userService.SearchUsers(userSearch)
	c.HTML(http.StatusOK, "user-search.html", gin.H{
		"title": "User Search", "User": user, "UserSearchList": userSearchList,
	})
}

func (h *AdminHandler) UserDetail(c *gin.Context) {
	user := services.GetUserSession(c)

	userid, _ := strconv.Atoi(c.Query("userid"))
	userDetail, _ := h.userService.GetUserById(userid)

	c.HTML(http.StatusOK, "user-detail.html", gin.H{
		"title": "User Detail", "User": user, "UserDetail": userDetail,
	})
}

func (h *AdminHandler) UserDetailPost(c *gin.Context) {
	user := services.GetUserSession(c)

	userid, _ := strconv.Atoi(c.Request.FormValue("userid"))
	action := c.Request.FormValue("btnSubmit")

	fmt.Println("action:", action)
	var message string
	alert := "alert-danger"

	if action == "Reset Password" {
		err := h.userService.ResetPassword(userid)
		if err != nil {
			message = "Error Reset Password"
		} else {
			message = "Success Reset Password"
			alert = "alert-success"
		}
	}

	userDetail, _ := h.userService.GetUserById(userid)

	c.HTML(http.StatusOK, "user-detail.html", gin.H{
		"title": "User Detail", "User": user, "UserDetail": userDetail, "message": message, "alert": alert,
	})
}

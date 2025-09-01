package handler

import (
	"booking-app/internal/entity"
	"booking-app/internal/services"
	"fmt"
	"net/http"

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
		"title": "User Management", "User": user,
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
		"title": "User Management", "User": user, "UserSearchList": userSearchList,
	})
}

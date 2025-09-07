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
	userService   services.UserServiceInterface
	lookupService services.LookupServiceInterface
}

func NewAdminHandler(userService services.UserServiceInterface, lookupService services.LookupServiceInterface) *AdminHandler {
	return &AdminHandler{
		userService:   userService,
		lookupService: lookupService,
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

	roles, _ := h.lookupService.GetAllActiveLookupByType(c, "USER_ROLE")
	status, _ := h.lookupService.GetAllActiveLookupByType(c, "USER_STATUS")

	userid, _ := strconv.Atoi(c.Query("userid"))
	userDetail, _ := h.userService.GetUserById(userid)

	c.HTML(http.StatusOK, "user-detail.html", gin.H{
		"title": "User Detail", "User": user, "UserDetail": userDetail,
		"Roles": roles, "Status": status,
	})
}

func (h *AdminHandler) UserDetailPost(c *gin.Context) {
	user := services.GetUserSession(c)

	roles, _ := h.lookupService.GetAllActiveLookupByType(c, "USER_ROLE")
	status, _ := h.lookupService.GetAllActiveLookupByType(c, "USER_STATUS")

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
	} else if action == "Create" {
		err := h.userService.CreateUser(c)
		if err != nil {
			message = "Error Create User"
		} else {
			message = "Success Create User"
			alert = "alert-success"
		}
	} else if action == "Update" {
		err := h.userService.UpdateUser(c)
		if err != nil {
			message = "Error Update User"
		} else {
			message = "Success Update User"
			alert = "alert-success"
		}
	}

	userDetail, _ := h.userService.GetUserById(userid)

	c.HTML(http.StatusOK, "user-detail.html", gin.H{
		"title": "User Detail", "User": user, "UserDetail": userDetail, "message": message, "alert": alert,
		"Roles": roles, "Status": status,
	})
}

func (h *AdminHandler) UserCreate(c *gin.Context) {
	user := services.GetUserSession(c)

	userDetail := entity.User{}
	roles, _ := h.lookupService.GetAllActiveLookupByType(c, "USER_ROLE")
	status, _ := h.lookupService.GetAllActiveLookupByType(c, "USER_STATUS")

	c.HTML(http.StatusOK, "user-detail.html", gin.H{
		"title": "User Detail", "User": user, "UserDetail": userDetail,
		"Roles": roles, "Status": status,
	})
}

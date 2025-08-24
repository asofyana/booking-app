package handler

import (
	"booking-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService services.UserServiceInterface
}

func NewAdminHandler(adminService services.UserServiceInterface) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

func (h *AdminHandler) UserSearch(c *gin.Context) {
	/*users, err := h.adminService.GetAllUsers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}*/

	user := services.GetUserSession(c)

	c.HTML(http.StatusOK, "user-search.html", gin.H{
		"title": "User Management", "User": user,
	})
}

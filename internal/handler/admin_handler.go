package handler

import (
	"booking-app/internal/entity"
	"booking-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService services.AdminServiceInterface
}

func NewAdminHandler(adminService services.AdminServiceInterface) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

func (h *AdminHandler) ShowAdminPage(c *gin.Context) {
	users, err := h.adminService.GetAllUsers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "admin.html", gin.H{
		"title": "Admin Dashboard",
		"users": users,
	})
}

func (h *AdminHandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBind(&user); err != nil {
		c.HTML(http.StatusBadRequest, "admin.html", gin.H{"error": err.Error()})
		return
	}

	if err := h.adminService.CreateUser(user); err != nil {
		c.HTML(http.StatusInternalServerError, "admin.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/admin")
}

func (h *AdminHandler) UpdateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBind(&user); err != nil {
		c.HTML(http.StatusBadRequest, "admin.html", gin.H{"error": err.Error()})
		return
	}

	if err := h.adminService.UpdateUser(user); err != nil {
		c.HTML(http.StatusInternalServerError, "admin.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/admin")
}

func (h *AdminHandler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	if err := h.adminService.DeleteUser(userId); err != nil {
		c.HTML(http.StatusInternalServerError, "admin.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/admin")
}

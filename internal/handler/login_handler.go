package handler

import (
	"booking-app/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService services.LoginServiceInterface
}

func NewLoginHandler(loginService services.LoginServiceInterface) *LoginHandler {
	return &LoginHandler{
		LoginService: loginService,
	}
}

func (s *LoginHandler) ShowLoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title": "Login",
		},
	)
}

func (s *LoginHandler) ProcessLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	err := s.LoginService.ProcessLogin(c, username, password)
	if err == nil {
		c.HTML(http.StatusOK, "home.html", nil)
		//c.Redirect(http.StatusTemporaryRedirect, "home")
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{"title": "Login", "message": "Invalid username or password"})

}

func (s *LoginHandler) ProcessLogout(c *gin.Context) {
	fmt.Println("ProcessLogout")
	services.InvalidateSession(c)
	c.Redirect(http.StatusTemporaryRedirect, "/")
	//c.HTML(http.StatusOK, "login.html", nil)
}

package services

import (
	"booking-app/internal/entity"
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret"))

func init() {
	store.Options.HttpOnly = true // since we are not accessing any cookies w/ JavaScript, set to true
	store.Options.Secure = true   // requires secuire HTTPS connection
	gob.Register(&entity.User{})
}

func Auth(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	sesUser, ok := session.Values["user"]
	fmt.Println("sesUser:", sesUser)
	if !ok {
		c.HTML(http.StatusForbidden, "login.html", nil)
		c.Abort()
		return
	}

	fmt.Println("middleware done")
	c.Next()
}

func BookingAuth(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	sesUser, ok := session.Values["user"]
	fmt.Println("sesUser:", sesUser)
	if !ok {
		c.HTML(http.StatusForbidden, "login.html", nil)
		c.Abort()
		return
	}
	c.Next()
}

func SetUserSession(c *gin.Context, user entity.User) {
	fmt.Println("Set User Session")
	session, _ := store.Get(c.Request, "session")
	session.Values["user"] = user
	session.Options.MaxAge = 600
	session.Save(c.Request, c.Writer)
}

func InvalidateSession(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
}

func GetUserSession(c *gin.Context) *entity.User {
	session, _ := store.Get(c.Request, "session")
	sesUser, ok := session.Values["user"]
	if ok {
		return sesUser.(*entity.User)
	}
	return &entity.User{}
}

package services

import (
	"booking-app/internal/entity"
	"encoding/gob"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret"))

func init() {
	store.Options.HttpOnly = true // since we are not accessing any cookies w/ JavaScript, set to true
	store.Options.Secure = false  // requires secuire HTTPS connection
	store.Options.Path = "/"
	store.Options.Domain = ""
	store.Options.SameSite = http.SameSiteStrictMode
	gob.Register(&entity.User{})
}

func Auth(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	_, ok := session.Values["user"]
	if !ok {
		c.HTML(http.StatusForbidden, "login.html", nil)
		c.Abort()
		return
	}

	c.Next()
}

func BookingAuth(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	_, ok := session.Values["user"]
	if !ok {
		c.HTML(http.StatusForbidden, "login.html", nil)
		c.Abort()
		return
	}
	c.Next()
}

func AdminAuth(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	_, ok := session.Values["user"]
	if !ok {
		c.HTML(http.StatusForbidden, "login.html", nil)
		c.Abort()
		return
	}

	user := GetUserSession(c)
	if user.Role != "admin" {
		c.HTML(http.StatusForbidden, "login.html", nil)
		c.Abort()
		return
	}

	c.Next()
}

func SetUserSession(c *gin.Context, user entity.User) {
	session, _ := store.Get(c.Request, "session")
	session.Values["user"] = user
	session.Options.MaxAge = 900
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

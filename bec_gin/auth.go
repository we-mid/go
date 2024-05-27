package bec_gin

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	adminUsername = os.Getenv("ADMIN_USERNAME")
	adminPassword = os.Getenv("ADMIN_PASSWORD")
)

func init() {
	assertBadPassword()
}

var CheckAuth gin.HandlerFunc = func(c *gin.Context) {
	if !IsLoggedIn(c) {
		c.JSON(401, errorRes{"Login required"})
		c.Abort()
	}
}

func AuthApis(r *gin.Engine, getConfigRes func() any) {
	r.POST("/api/login", ratelimitLoginApi(), func(c *gin.Context) {
		json := &loginReq{}
		_ = c.BindJSON(json)
		if json.Username == adminUsername && json.Password == adminPassword {
			setUser(c, json.Username)
			c.JSON(200, successRes{})
		} else {
			c.JSON(400, errorRes{"Wrong username or password"})
			SecurityLog(c, "Failed attempt to login")
		}
	})
	r.GET("/api/session", CheckAuth, func(c *gin.Context) {
		c.JSON(200, getConfigRes())
	})
}

func IsLoggedIn(c *gin.Context) bool {
	_, ok := getUser(c)
	return ok
}

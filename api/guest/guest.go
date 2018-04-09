package guest

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"seed/o/auth"
	"seed/o/user"
)

type GuestServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewGuestServer(parent *gin.RouterGroup, name string) *GuestServer {
	var s = GuestServer{
		RouterGroup: parent.Group(name),
	}
	s.POST("register", s.register)
	return &s
}

func (s *GuestServer) register(c *gin.Context) {
	var u *user.User
	c.BindJSON(&u)
	u.Role = user.GUEST
	web.AssertNil(u.Create())
	var auth = auth.Create(u.ID, u.Role)
	s.SendData(c, map[string]interface{}{
		"token":     auth.ID,
		"user_info": u,
	})
	s.SendData(c, u)
}

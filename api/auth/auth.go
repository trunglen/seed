package auth

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"seed/o/auth"
	"seed/o/user"
)

type AuthServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewAuthServer(parent *gin.RouterGroup, name string) *AuthServer {
	var s = AuthServer{
		RouterGroup: parent.Group(name),
	}
	s.POST("login", s.login)
	s.GET("super-admin", s.checkSuperAdmin)
	return &s
}

func (s *AuthServer) login(c *gin.Context) {
	var loginInfo = struct {
		Uname    string
		Password string
	}{}
	c.BindJSON(&loginInfo)
	user, err := user.Login(loginInfo.Uname, loginInfo.Password)
	web.AssertNil(err)
	var auth = auth.Create(user.ID, user.Role)
	s.SendData(c, map[string]interface{}{
		"token":     auth.ID,
		"user_info": user,
	})
}

func (s *AuthServer) checkSuperAdmin(c *gin.Context) {
	var superAdmin, err = user.GetAdmin("", "super-admin")
	web.AssertNil(err)
	s.SendData(c, superAdmin)
}

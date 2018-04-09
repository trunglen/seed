package user

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"seed/o/user"
)

type UserServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewUserServer(parent *gin.RouterGroup, name string) *UserServer {
	var s = UserServer{
		RouterGroup: parent.Group(name),
	}
	s.GET("list", s.getUsers)
	s.POST("create", s.createUser)
	s.GET("delete", s.deleteUser)
	return &s
}

func (s *UserServer) getUsers(c *gin.Context) {
	var users, err = user.GetUsers()
	web.AssertNil(err)
	s.SendData(c, users)
}

func (s *UserServer) createUser(c *gin.Context) {
	var user *user.User
	c.BindJSON(&user)
	web.AssertNil(user.Create())
	s.SendData(c, user)
}

func (s *UserServer) deleteUser(c *gin.Context) {
	var id = c.Query("id")
	web.AssertNil(user.DeleteUserByID(id))
	s.Success(c)
}

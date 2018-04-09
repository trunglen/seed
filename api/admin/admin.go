package admin

import (
	"g/x/web"
	"github.com/gin-gonic/gin"
	"seed/api/admin/auth"
	"seed/api/admin/category"
	"seed/api/admin/post"
	"seed/api/admin/user"
)

type AdminServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewAdminServer(parent *gin.RouterGroup, name string) *AdminServer {
	var s = AdminServer{
		RouterGroup: parent.Group(name),
	}
	post.NewPostServer(s.RouterGroup, "post")
	category.NewCategoryServer(s.RouterGroup, "category")
	user.NewUserServer(s.RouterGroup, "user")
	auth.NewAuthServer(s.RouterGroup, "auth")
	return &s
}

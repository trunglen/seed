package admin

import (
	"github.com/gin-gonic/gin"
	"seed/api/admin/category"
	"seed/api/admin/post"
	"seed/api/admin/user"
	"seed/g/x/web"
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
	return &s
}

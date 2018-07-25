package admin

import (
	"github.com/gin-gonic/gin"
	"rocky-springs-86767/api/admin/category"
	"rocky-springs-86767/api/admin/post"
	"rocky-springs-86767/api/admin/user"
	"rocky-springs-86767/g/x/web"
	"rocky-springs-86767/middleware"
)

type AdminServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewAdminServer(parent *gin.RouterGroup, name string) *AdminServer {
	var s = AdminServer{
		RouterGroup: parent.Group(name),
	}
	s.Use(middleware.MustBeAdmin)
	post.NewPostServer(s.RouterGroup, "post")
	category.NewCategoryServer(s.RouterGroup, "category")
	user.NewUserServer(s.RouterGroup, "user")
	return &s
}

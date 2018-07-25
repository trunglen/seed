package api

import (
	"rocky-springs-86767/api/admin"
	"rocky-springs-86767/api/auth"
	"rocky-springs-86767/api/guest"
	"rocky-springs-86767/api/public"

	"github.com/gin-gonic/gin"
)

func NewApiServer(root *gin.RouterGroup) {
	admin.NewAdminServer(root, "admin")
	auth.NewAuthServer(root, "auth")
	public.NewPublicServer(root, "public")
	guest.NewGuestServer(root, "guest")
}

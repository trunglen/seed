package api

import (
	"seed/api/admin"
	"seed/api/guest"
	"seed/api/public"

	"github.com/gin-gonic/gin"
)

func NewApiServer(root *gin.RouterGroup) {
	admin.NewAdminServer(root, "admin")
	public.NewPublicServer(root, "public")
	guest.NewGuestServer(root, "guest")
}

package router

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/forever1young/account-onboard-alltoken/config"
	"github.com/forever1young/account-onboard-alltoken/controller"
	"github.com/forever1young/account-onboard-alltoken/middleware"
)

func New(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/logout", controller.Logout)

	admin := r.Group("/api/admin", middleware.RequireOperatorSession())
	{
		admin.POST("/oauth/link", controller.GenerateAuthLink)
		admin.POST("/oauth/callback", controller.HandleCallback)
		admin.GET("/accounts", controller.ListAccounts)
		admin.PUT("/accounts/:id", controller.UpdateAccount)
	}

	export := r.Group("/api/export", middleware.RequireExportAuth(cfg.ExportAPIKey, splitIPs(cfg.ExportAPIAllowedIPs)))
	{
		export.GET("/accounts", controller.PullAccounts)
		export.POST("/ack", controller.AckDelivery)
		export.POST("/report", controller.ReportStatus)
	}

	return r
}

func splitIPs(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}

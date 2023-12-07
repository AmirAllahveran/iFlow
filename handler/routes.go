package handler

import (
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
	"github.com/smf8/http-monitor/common"
	"github.com/smf8/http-monitor/middleware"
)

// RegisterRoutes registers routes with their corresponding handler function
// functions are defined in handler package
func (h *Handler) RegisterRoutes(v *echo.Group) {

	v.Use(middleware.JWT(common.JWTSecret))
	v.Use(middleware2.RemoveTrailingSlash())

	// adding white list
	middleware.AddToWhiteList("/api/users/login", "POST")
	middleware.AddToWhiteList("/api/users", "POST")
	middleware.AddToWhiteList("/api/experiment", "POST")

	userGroup := v.Group("/users")
	userGroup.POST("", h.SignUp)
	userGroup.POST("/login", h.Login)

	//moduleGroup := v.Group("/module")
	//moduleGroup.GET("", h.CreateModule)

	experimentGroup := v.Group("/experiment")
	experimentGroup.POST("", h.RunExperiment)

}

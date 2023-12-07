package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/smf8/http-monitor/docker"
	"github.com/smf8/http-monitor/store"
)

// require validator to add "required" tag to every struct field in the package
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Handler struct {
	st  *store.Store
	dck *docker.Client
}

// NewHandler creates a new handler with given store instance
func NewHandler(st *store.Store, dck *docker.Client) *Handler {
	return &Handler{st: st, dck: dck}
}

func extractID(c echo.Context) uint {
	e := c.Get("user").(*jwt.Token)
	claims := e.Claims.(jwt.MapClaims)
	id := uint(claims["id"].(float64))
	return id
}

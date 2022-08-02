package handler

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) http.Handler {
	handler := &Server{db: db}
	e := echo.New()
	e.POST("/link", handler.createLink)
	return e
}

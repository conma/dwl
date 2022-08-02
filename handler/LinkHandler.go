package handler

import (
	"dwl/entity"
	"dwl/request"
	"dwl/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strconv"
)

func (s Server) createLink(context echo.Context) (err error) {
	req := new(request.CreateLink)
	if err = context.Bind(req); err != nil {
		return
	}

	existing := entity.Link{}
	s.db.Debug().Where("id = ? and is_deleted = ?", req.Link, false).Find(&existing)
	if existing.Id != nil && *existing.CurrentUsr == req.User {
		res := response.CreateLinkResponse{
			Link: existing.Id,
			User: existing.CurrentUsr,
		}
		return context.JSONPretty(http.StatusOK, res, "  ")
	}

	vFalse := false
	link := entity.Link{
		Id:         &req.Link,
		IsDeleted:  &vFalse,
		CurrentUsr: &req.User,
	}

	if existing.Id != nil && *existing.CurrentUsr != req.User {
		randLink := req.Link + strconv.Itoa(rand.Intn(9999))
		for existLink(randLink, s.db) {
			randLink = req.Link + strconv.Itoa(rand.Intn(9999))
		}
		link.Id = &randLink
	}

	s.db.Create(&link)
	res := response.CreateLinkResponse{
		Link: link.Id,
		User: link.CurrentUsr,
	}
	return context.JSONPretty(http.StatusOK, res, "  ")
}

func existLink(link string, db *gorm.DB) bool {
	currentLink := entity.Link{}
	db.Debug().Where("id = ? and is_deleted = false", link).Find(&currentLink)
	return currentLink.Id != nil
}

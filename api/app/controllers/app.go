package controllers

import (
	"api/app"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
	DB *gorm.DB
}

func (c *App) Before() revel.Result {
	c.DB = app.Gorm
	return nil
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) BadRequestJSON(s interface{}) revel.Result {
	c.Response.Status = http.StatusBadRequest
	return c.RenderJSON(s)
}

func (c App) ErrorJSON(s interface{}) revel.Result {
	c.Response.Status = http.StatusInternalServerError
	return c.RenderJSON(s)
}

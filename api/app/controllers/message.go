package controllers

import (
	"log"

	"api/app/models"

	"github.com/revel/revel"
)

type Message struct {
	App
}

func (c *Message) List() revel.Result {
	messages, err := models.Messages(c.DB)
	if err != nil {
		log.Println(err)
		return c.ErrorJSON(Error{"Message取得に失敗しました"})
	}

	return c.RenderJSON(messages)
}

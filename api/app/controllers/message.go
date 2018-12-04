package controllers

import (
	"github.com/revel/revel"
)

type Message struct {
	App
}

func (c Message) List(message string) revel.Result {
	if message == "" {
		return c.BadRequestJSON(Result{"message is empty"})
	}

	return c.RenderJSON(nil)
}

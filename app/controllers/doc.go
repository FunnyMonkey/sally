package controllers

import (
	"github.com/robfig/revel"
)

type Doc struct {
	App
}

func (c Doc) Index() revel.Result {
	return c.Render()
}

func (c Doc) Create() revel.Result {
	return c.Render()
}

func (c Doc) Show() revel.Result {
	return c.Render()
}

func (c Doc) Update() revel.Result {
	return c.Render()
}

func (c Doc) Destroy() revel.Result {
	return c.Render()
}

package controllers

import (
	"github.com/robfig/revel"
)

type Text struct {
	App
}

func (c Text) Index() revel.Result {
	return c.Render()
}

func (c Text) Create() revel.Result {
	return c.Render()
}

func (c Text) Show() revel.Result {
	return c.Render()
}

func (c Text) Update() revel.Result {
	return c.Render()
}

func (c Text) Destroy() revel.Result {
	return c.Render()
}

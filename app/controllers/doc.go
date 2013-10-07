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

func (c Doc) Copy() revel.Result {
	return c.Render()
}

// Doc routes will just redirect to the corresponding type/:id path.
// This way we can have generic doc routes. Additionally docs will be stored in
// both the doctype

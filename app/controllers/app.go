package controllers

import (
	"github.com/jgraham909/revmgo"
	"github.com/robfig/revel"
)

type App struct {
	revmgo.MongoController
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

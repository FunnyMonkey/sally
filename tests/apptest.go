package tests

import "github.com/robfig/revel"

type AppTest struct {
	revel.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t AppTest) TestDoNothing() {
	// Do nothing for now
}

func (t *AppTest) After() {
	println("Tear down")
}

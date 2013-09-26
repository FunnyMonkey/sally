package tests

import "github.com/robfig/revel"

type DocTest struct {
	revel.TestSuite
}

func (t *DocTest) Before() {
	println("Set up")
}

func (t *DocTest) TestValidateEmptyTitle() {

}

func (t *DocTest) TestValidateEmptyType() {

}

func (t *DocTest) TestSave() {
}

func (t *DocTest) TestFind() {
}

func (t *DocTest) TestDelete() {
}

// TEST API endpoints here
//
//GET      /doc              Doc.Index
//GET      /doc/doc_id       Doc.Show

func (t *DocTest) After() {
	println("Tear down")
}

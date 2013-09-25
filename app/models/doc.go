package models

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"reflect"
	"sally/app"
	"strings"
)

// All sally content types should embed the base Doc struct
type Doc struct {
	Id        bson.ObjectId   `bson:"_id,omitempty"`
	Ancestors []bson.ObjectId `bson: ancestors`
	History   []bson.ObjectId `bson: history`
	License   string          `bson:license`
	Type      string          `bson:type`
	Title     string          `bson:title`
}

func Collection(m interface{}, s *mgo.Session) *mgo.Collection {
	typ := reflect.TypeOf(m)
	i := strings.LastIndex(typ.String(), ".") + 1
	n := typ.String()[i:len(typ.String())]

	var found bool
	var c string
	if c, found = revel.Config.String("sally.db.collection." + n); !found {
		c = n
	}
	return s.DB(app.DB).C(c)
}

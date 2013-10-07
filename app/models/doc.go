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
	DocId     bson.ObjectId   `bson:"_id,omitempty"` // Unique document ID
	Ancestors []bson.ObjectId `bson: ancestors`      // Stores document IDs this was copied from (parent branches/forks)
	History   []bson.ObjectId `bson: history`        // Stores document type IDs document content & revisions
	License   string          `bson:license`         // CC license
	Type      string          `bson:type`            // Specialty document type
	Title     string          `bson:title`           // String of document (latest version)
}

// Does the necessary housekeeping to fork a document.
func (d *Doc) PreCopy() {
	d.Ancestors = append(d.Ancestors, d.DocId)
	d.DocId = bson.NewObjectId()
}

// Does the necessary housekeeping to save a revision of a document.
// Returns the ObjectId for the new document *type* id.
//
// It would be nice to refactor this somehow but not require each
// document type to implement it.
func (d *Doc) PreSave(cur bson.ObjectId) bson.ObjectId {
	d.History = append(d.History, cur)
	return bson.NewObjectId()
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

func DocCollection(s *mgo.Session) *mgo.Collection {
	var found bool
	var c string
	if c, found = revel.Config.String("sally.db.collection.doc"); !found {
		c = "doc"
	}
	return s.DB(app.DB).C(c)
}

func LoadDocByObjectId(s *mgo.Session, DocId bson.ObjectId) *Doc {
	d := new(Doc)

	query := DocCollection(s).FindId(DocId)
	query.One(d)
	return d
}

func LoadDocById(s *mgo.Session, Id string) *Doc {
	DocId := bson.ObjectIdHex(Id)
	return GetArticleByObjectId(s, DocId)
}

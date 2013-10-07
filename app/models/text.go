package models

// We just have a value string which represents the text of this document.
type Text struct {
	Doc       Doc
	DocTypeId bson.ObjectId `bson:"_id,omitempty"` // Unique document type ID
	Value     string        `bson:value`
}

func (t *Text) Save(s *mgo.Session) error {
	c := Collection(t, s)
	_, err := c.Upsert(bson.M{"_id": t.DocTypeId}, t)
	if err != nil {
		revel.WARN.Printf("Unable to save text document: %v error %v", t, err)
	}
	return err
}

func GetTextByObjectId(s *mgo.Session, Id bson.ObjectId) *Text {
	t := new(Text)

	query := Collection(t, s).FindId(Id)
	query.One(t)
	return t
}

func (t *Text) Delete(s *mgo.Session) error {
	c := Collection(t, s)
	err := c.RemoveId(t.DocTypeId)
	if err != nil {
		revel.WARN.Printf("Undable to delete text document: %v error %v", t, err)
	}
	return err
}

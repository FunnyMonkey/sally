package models

// All sally content types should embed the base Doc struct
type Text struct {
	Doc   Doc
	Value string `bson:value`
}

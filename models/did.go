package models

import (
	"github.com/satori/go.uuid"
	"fmt"
)

type Did struct {
	ID        uuid.UUID     `db:"id" json:"id"`
	UserID    uuid.UUID     `db:"user_id" json:"user_id"`
	Text      string      	`db:"text" json:"text"`
	Source	  Source     	`db:"source" json:"source"`
	Tags	  []Tag		`db:"tags" json:"tags"`
	CreatedAt string        `db:"created_at" json:"created_at"`
}

func (d *Did) fullTextSearchFilter() string {
	return ` WHERE searchable @@ plainto_tsquery($1)`
}

func (d *Did) GetDids(userID uuid.UUID,q string) DidCollection {

	// using make will marshal to an empty array when no results
	dids := make([]Did, 0)
	collection := DidCollection{Dids:dids}

	err :=  DB.SelectDoc("id", "user_id", "text, created_at").
		From("dids").
		Where("user_id = $1", userID).
		Scope(d.fullTextSearchFilter(), q).
		Many("tags", DB.
			Select("id", "user_id", "text", "created_at").
			From("tags").
			Where("tags.id = (SELECT tag_id FROM did_tag WHERE did_id = dids.id)")).
		One("source", DB.
			Select("id", "name").
			From("sources").
			Where("dids.source_id = id")).
		Limit(20).
		QueryStructs(&collection.Dids)
	if err != nil {
		fmt.Println(err)
	}

	return collection
}




package models

type DidCollection struct {
	Dids []Did
}

func (dc *DidCollection) LastId() *string {
	if len(dc.Dids) == 0 {
		return nil
	}

	lastId := dc.Dids[len(dc.Dids)-1].ID.String()
	return &lastId
}

func (dc *DidCollection) Len() int {
	return len(dc.Dids)
}
package model

type Tag struct {
	ID1    string `json:"id1" xorm:"varchar(6)"`
	ID2    string `json:"id2" xorm:"varchar(6)"`
	ID3    string `json:"id3" xorm:"varchar(6)"`
	Height int    `json:"height" xorm:"int"`
	Name   string `json:"name" xorm:"varchar(45) not null"`
}

func (t *Tag) TableName() string {
	return "tags"
}

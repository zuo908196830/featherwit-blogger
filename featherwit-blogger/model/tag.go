package model

// 使用三个id和height唯一确定一个标签，规则如下
// 标签使用树形方式管理，height表示该标签在树上的深度，三个ID分别表示其在对应深度上从属的节点，使用height和对应高度的ID唯一确定一个标签
type Tag struct {
	Id1    string `json:"id1" xorm:"varchar(6)"`
	Id2    string `json:"id2" xorm:"varchar(6)"`
	Id3    string `json:"id3" xorm:"varchar(6)"`
	Height int    `json:"height" xorm:"int"`
	Name   string `json:"name" xorm:"varchar(45) not null"`
}

func (t *Tag) TableName() string {
	return "tags"
}

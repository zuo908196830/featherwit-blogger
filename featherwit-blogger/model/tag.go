package model

// 使用三个id和height唯一确定一个标签，规则如下
// 标签使用树形方式管理，height表示该标签在树上的深度，三个ID分别表示其在对应深度上从属的节点，使用height和对应高度的ID唯一确定一个标签
type Tag struct {
	ID          int64  `json:"id" xorm:"bigint not null pk id autoincr"`
	Name        string `json:"name" xorm:"varchar(45) not null"`
	SearchCount int64  `json:"searchCount" xorm:"bigint not null default 0"`
	Charset     string `xorm:"'ENGINE=InnoDB CHARSET=utf8'"`
}

func (t *Tag) TableName() string {
	return "tags"
}

package model

type Blob struct {
	BaseModel
	UserId      int    `json:"userId" xorm:"int"`
	Title       string `json:"title" xorm:"text"`
	Content     string `json:"content" xrom:"longtext"`
	Views       int64  `json:"views" xorm:"bigint"` //浏览量
	CommonCount int64  `json:"commonCount" xorm:"bigint"`
	LikeCount   int64  `json:"likeCount" xorm:"bigint"`
}

func (b *Blob) TableName() string {
	return "blob"
}

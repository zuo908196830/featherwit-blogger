package model

import "time"

type ShopType struct {
	Id int64 `json:"id" xorm:"bigint(20) pk not null"`
	Name string `json:"name" xorm:"varchar(32)"`
	Icon string `json:"icon" xomr:"varchar(255)"`
	Sort int `json:"sort" xorm:"int(3)"`
	CreateAt time.Time `json:"createAt" xorm:"-"`
	UpdateAt time.Time `json:"updateAt" xorm:"-"`
}

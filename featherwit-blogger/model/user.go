package model

type User struct {
	BaseModel
	Username  string `json:"userName" xorm:"varchar(25) 'username'"`
	Password  string `json:"password" xorm:"varchar(45) 'password'"`
	Role      int    `json:"role" xorm:"int"`
	Nickname  string `json:"nickname" xorm:"varchar(45)"`
	Telephone string `json:"telephone" xorm:"varchar(20)"`
	Headshot  string `json:"headshot" xorm:"varchar(50)"`
}

func (u *User) TableName() string {
	return "users"
}

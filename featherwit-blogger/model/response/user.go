package response

type Login struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
}

type UserData struct {
	Username  string `json:"username" xorm:"varchar(25) 'username' index"`
	Role      int    `json:"role" xorm:"int"`
	Nickname  string `json:"nickname" xorm:"varchar(45)"`
	Telephone string `json:"telephone" xorm:"varchar(20)"`
	Mail      string `json:"mail" xorm:"varchar(45)"`
	Profile   string `json:"profile" xorm:"varchar(400)"`
}

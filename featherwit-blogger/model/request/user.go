package request

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Role      int    `json:"role" binding:"required"`
	Nickname  string `json:"nickname"`
	Mail      string `json:"mail"`
	Telephone string `json:"telephone"`
	Profile   string `json:"profile"`
}

type AddAttentionRequest struct {
	AUsername string `json:"aUsername" uri:"aUsername"`
}

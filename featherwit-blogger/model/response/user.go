package response

type Login struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
}

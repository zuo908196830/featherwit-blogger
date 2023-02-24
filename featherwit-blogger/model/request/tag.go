package request

type AddTagRequest struct {
	Tags []*TagTree `json:"tags"`
}

type TagTree struct {
	Name     string     `json:"name"`
	Children []*TagTree `json:"children"`
}

package structs

type Post struct {
	Images []string `json:"images" validate:"required"`
	Text   string   `json:"text" validate:"required"`
}

type Comments struct {
	Comment string `json:"comment"`
	PostId  int    `json:"postId"`
}

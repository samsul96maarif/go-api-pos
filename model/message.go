package model

type Message struct {
	BaseModel
	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content"`
	IsRead  bool   `json:"is_read"`
}

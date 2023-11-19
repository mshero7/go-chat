package main




type Message struct {
	Sender   string `json:"sender,omitempty"`
	Recipent string `json:"recipient,omitempty"`
	Content  string `json:"content,omitempty"`
}


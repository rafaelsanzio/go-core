package handlers

type OkPayload struct {
	Health int    `json:"health,omitempty"`
	Test   string `json:"test,omitempty"`
}

type UserEntityPayload struct {
	Name string
	Age  string
}

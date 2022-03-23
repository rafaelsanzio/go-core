package handlers

type OkPayload struct {
	Health int    `json:"health,omitempty"`
	Test   string `json:"test,omitempty"`
}

type UserEntityPayload struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
}

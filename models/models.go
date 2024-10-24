package models

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type EmailRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type Payload struct {
	Ref string `json:"ref"`
}

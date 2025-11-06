package models

type Album struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

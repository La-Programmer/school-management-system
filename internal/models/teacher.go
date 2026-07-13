package models

type Teacher struct {
	Id        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Class     string `json:"class,omitempty"`
	Subject   string `json:"subject,omitempty"`
}

type MultiTeacherResp struct {
	Status string    `json:"status"`
	Count  int       `json:"count"`
	Data   []Teacher `json:"data"`
}

package models

type Program struct {
	Uid  string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

package models

type Social struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Title      string `json:"title"`
	ShowLink   string `json:"show_link"`
	ActualLink string `json:"actual_link"`
}

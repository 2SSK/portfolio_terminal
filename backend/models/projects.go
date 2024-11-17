package models

type Project struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

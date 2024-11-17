package models

type Skill struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Type string `json:"type"`
}

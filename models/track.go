package models

type Track struct {
	ID 	   uint   `json:"id" gorm:"primary key"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

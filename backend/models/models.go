package models

type Todo struct {
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title string `json:"title"`
}

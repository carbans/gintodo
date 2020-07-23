package models

type Task struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Owner string `json:"owner"`
  }
package src

import (
	"time"
)

type TodoItem struct {
	Id int `json:"id"`
	Subject string `json:"subject" binding:"max=30"`
	Status int `json:"status" binding:"min=0,max=1"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

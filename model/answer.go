package model

import "github.com/jinzhu/gorm"

var answers map[string]string

type Answer struct {
	gorm.Model
	userID   int64
	Question string
	Answer   string
}

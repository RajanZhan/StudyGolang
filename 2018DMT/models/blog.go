package models

import "time"

//博客
type Blog struct {
	Id          int       `gorm:"primary_key" json:"Id"`
	PublisherId int       `json:"PublisherId"`
	Time        time.Time `json:"Time"`
	Status      int       `json:"Status"`
	Type        int       `json:"Type"`
	Title       string    `json:"Title"`
	Context     string    `gorm:"type:varchar(10239)" json:"Context"`
	Readed      int       `json:"Readed"`
}
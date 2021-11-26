package models

type Role struct {
	Id          int    `gorm:"column:id" json:"id"`
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	Status      int    `gorm:"column:status" json:"status"`
	AddTime     int    `gorm:"column:add_time" json:"addTime"`
}

func (Role) TableName() string {
	return "role"
}
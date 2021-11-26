package models

type Focus struct {
	Id        int    `form:"id" json:"id"`
	Title     string `form:"title" json:"title"`
	FocusType int    `form:"focus_type" json:"focusType"`
	FocusImg  string `form:"focusImg" json:"focusImg"`
	Link      string `form:"link" json:"link"`
	Sort      int    `form:"sort" json:"sort"`
	Status    int    `form:"status" json:"status"`
	AddTime   int    `form:"addTime" json:"addTime"`
}

func (Focus) TableName() string {
	return "focus"
}

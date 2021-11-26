package models

type Manager struct {
	Id       int    `gorm:"column:id" json:"id" form:"id"`
	Username string `gorm:"column:username" json:"username" form:"username"`
	Password string `gorm:"column:password" json:"password" form:"password"`
	Mobile   string `gorm:"column:mobile" json:"mobile" form:"mobile"`
	Email    string `gorm:"column:email" json:"email" form:"email"`
	Status   int    `gorm:"column:status" json:"status" form:"status"`
	RoleId   int    `gorm:"column:role_id" json:"roleId" form:"roleId"`
	AddTime  int    `gorm:"column:add_time" json:"addTime" form:"addTime"`
	IsSuper  int    `gorm:"column:is_super" json:"isSuper" form:"isSuper"`
	Role     Role   `gorm:"foreignKey:RoleId;references:Id"`
}

func (Manager) TableName() string {
	return "manager"
}

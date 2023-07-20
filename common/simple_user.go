package common

type SimpleUser struct {
	SQLModel  `json:",inline"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	Role      string `json:"role" gorm:"column:role"`
	Avatar    *Image `json:"avatar" gorm:"column:avatar"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdminOrUser bool) {
	u.GenUID(DbTypeUser)
}

package models

type Users struct {
	Id        int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"` // 主键
	Username  string `gorm:"column:username;type:varchar(255)" json:"username"`
	Password  string `gorm:"column:password;type:varchar(255)" json:"password"`
	CreatedAt int64  `gorm:"column:created_at;type:int(11)" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;type:int(11)" json:"updated_at"`
}

func (m *Users) TableName() string {
	return "users"
}

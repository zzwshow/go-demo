package models

type UserModel struct {
	UserID	uint64 `gorm:"primary_key"`
	Name string
	Age int
}

func (UserModel) TableName() string {
	return "user"
}


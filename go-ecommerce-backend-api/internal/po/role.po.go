package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type:int; not null; primaryKey; autoIncrement; comment:'primary Key id ID'"`
	RoleName string `gorm:"column:role_name; not null;"`
	RoleNote string `gorm:"column:role_note; type:text;"`
}

func (r *Role) TableName() string {
	return "go_db_roles"
}

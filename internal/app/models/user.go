package models

import (
	"database/sql/driver"

	"messenger-server/internal/pkg/database"

	"github.com/enorith/authenticate"
	"github.com/enorith/http/content"
)

type Password []byte

// Value returns a driver Value.
// Value must not panic.
func (p Password) Value() (driver.Value, error) {
	return []byte(p), nil
}

type User struct {
	content.Request `gorm:"-" json:"-"`
	ID              int64    `gorm:"column:id;primaryKey;not null;type:int;autoIncrement" json:"id"`
	Name            string   `gorm:"column:name;type:varchar(255)" json:"name" input:"name" validate:"required"`
	Username        string   `gorm:"column:username;uniqueIndex:idx_unique_username;type:varchar(128)" json:"username" input:"username" validate:"required|unique:messenger_users"`
	Password        Password `gorm:"column:password;type:varchar(255)" json:"-" input:"password" validate:"required"`
	database.WithTimestamps
}

func (u User) UserIdentifier() authenticate.UserIdentifier {
	return authenticate.Identifier(u.ID)
}

func (u User) TableName() string {
	return "messenger_users"
}

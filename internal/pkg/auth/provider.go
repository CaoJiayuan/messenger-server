package auth

import (
	"fmt"

	"github.com/enorith/authenticate"
	"messenger-server/internal/app/models"
	"github.com/enorith/framework/authentication"
	"github.com/enorith/http/contracts"
	"gorm.io/gorm"
)

var (
	UsernameField   = "username"
	PasswordField   = "password"
	AuthFailedError error
)

type UserProvider struct {
	DB *gorm.DB
}

func (up *UserProvider) FindUserById(id authenticate.UserIdentifier) (authenticate.User, error) {
	var user models.User
	e := up.DB.First(&user, id.Int64()).Error
	return user, e
}

func (up *UserProvider) Attempt(r contracts.RequestContract) (authenticate.User, error) {
	var user models.User
	e := up.DB.Where(fmt.Sprintf("%s = ?", UsernameField), r.GetString(UsernameField)).Find(&user).Error
	if e != nil {
		return nil, e
	}
	if user.ID == 0 {
		return nil, AuthFailedError
	}
	if !authentication.Compare(user.Password, r.Get(PasswordField)) {
		return user, AuthFailedError
	}
	return user, e
}

func NewUserProvider(db *gorm.DB) *UserProvider {
	return &UserProvider{DB: db}
}

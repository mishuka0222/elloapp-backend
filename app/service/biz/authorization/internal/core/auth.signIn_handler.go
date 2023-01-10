package core

import (
	"errors"
	"fmt"
	"github.com/gogo/protobuf/types"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthSingIN
// TODO: need to write logic
func (c *AuthorizationCore) AuthSingIN(in *authorization.AuthSignInRequest) (*types.Empty, error) {
	var user models.Users
	if err := c.svcCtx.Gorm.Where("username = ?", in.Username).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		err = fmt.Errorf("can not get users record (%v)", err)
		c.Logger.Error(err)
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		err = fmt.Errorf("no such user exists with username = %s", in.Username)
		c.Logger.Error(err)
		return nil, err
	}

	var userEllo models.UsersEllo
	if err := c.svcCtx.Gorm.Where("user_id = ?", user.ID).First(&userEllo).Error; err != nil && err != gorm.ErrRecordNotFound {
		err = fmt.Errorf("can not get users_ello record (%v)", err)
		c.Logger.Error(err)
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		err = errors.New("no such users_ello exists")
		c.Logger.Error(err)
		return nil, err
	} else if !userEllo.EmailConfirmed {
		err := fmt.Errorf("email not confirmed for user %s", user.Username)
		c.Logger.Error(err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userEllo.Password), []byte(in.Password)); err != nil {
		err = errors.New("user password is incorrect")
		c.Logger.Error(err)
		return nil, err
	}

	return nil, errors.New(" Unimplemented")
}

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
	var userEllo models.UsersEllo
	if err := c.svcCtx.Gorm.Where("username = ?", in.Username).First(&userEllo).Error; err != nil && err != gorm.ErrRecordNotFound {
		err = fmt.Errorf("can not get users_ello record (%v)", err)
		c.Logger.Error(err)
		return nil, err
	} else if err == gorm.ErrRecordNotFound {
		err = errors.New("username or password incorrect")
		c.Logger.Error(err)
		return nil, err
	} else if !userEllo.EmailConfirmed {
		err := fmt.Errorf("email not confirmed for user %s", userEllo.Username)
		c.Logger.Error(err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userEllo.Password), []byte(in.Password)); err != nil {
		err = errors.New("username or password are incorrect")
		c.Logger.Error(err)
		return nil, err
	}

	return &types.Empty{}, nil
}

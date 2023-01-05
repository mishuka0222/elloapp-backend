package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/internal/config"
)

type Dao struct {
}

func New(c config.Config) *Dao {
	return &Dao{}
}

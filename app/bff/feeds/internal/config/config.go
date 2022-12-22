package config

import "github.com/teamgram/marmota/pkg/stores/sqlx"

type Config struct {
	Mysql sqlx.Config
}

package dao

import (
	"flag"
	"github.com/oschwald/geoip2-golang"
	sync_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/client"
	authsession_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/internal/dao/dao/mysql_dao"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

var (
	mmdb string
)

func init() {
	flag.StringVar(&mmdb, "mmdb", "./GeoLite2-City.mmdb", "mmdb")
}

type Dao struct {
	*mysql_dao.Mysql
	authsession_client.AuthsessionClient
	MMDB *geoip2.Reader
	sync_client.SyncClient
}

func New(c config.Config) *Dao {
	db := sqlx.NewMySQL(&c.Mysql)
	MMDB, err := geoip2.Open(mmdb)
	if err != nil {
		// panic(err)
	}
	return &Dao{
		Mysql:             mysql_dao.NewMysqlDao(db),
		MMDB:              MMDB,
		AuthsessionClient: authsession_client.NewAuthsessionClient(rpcx.GetCachedRpcClient(c.AuthsessionClient)),
		SyncClient:        sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.SyncClient)),
	}
}

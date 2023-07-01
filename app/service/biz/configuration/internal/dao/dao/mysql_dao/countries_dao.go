package mysql_dao

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/configuration/internal/dao/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

var _ *sql.Result

type CountriesDAO struct {
	db *sqlx.DB
}

func NewCountriesDAO(db *sqlx.DB) *CountriesDAO {
	return &CountriesDAO{db}
}

// SelectList
// select id, code, "name", flag  from countries_flags
func (dao *CountriesDAO) SelectList(ctx context.Context) (rList []dataobject.CountriesDO, err error) {
	var (
		query  = `select id, code, name, flag  from countries_flags`
		values []dataobject.CountriesDO
	)
	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	rList = values

	return
}

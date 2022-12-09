package mysql_dao

import (
	"context"
	"database/sql"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/bff/configuration/internal/dao/dataobject"
	"github.com/zeromicro/go-zero/core/logx"
)

var _ *sql.Result

type CountriesDAO struct {
	db *sqlx.DB
}

func NewCountriesDAO(db *sqlx.DB) *CountriesDAO {
	return &CountriesDAO{db}
}

// SelectList
// select id, code, "name", flag from countries
func (dao *CountriesDAO) SelectList(ctx context.Context) (rList []dataobject.CountriesDO, err error) {
	var (
		query  = `select id, code, "name", flag from countries`
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

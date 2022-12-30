package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/internal/dal/dao/mysql_dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.DocumentsDAO
	*mysql_dao.EncryptedFilesDAO
	*mysql_dao.PhotosDAO
	*mysql_dao.PhotoSizesDAO
	*mysql_dao.VideoSizesDAO
	*sqlx.CommonDAO
}

func newMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:                db,
		DocumentsDAO:      mysql_dao.NewDocumentsDAO(db),
		EncryptedFilesDAO: mysql_dao.NewEncryptedFilesDAO(db),
		PhotosDAO:         mysql_dao.NewPhotosDAO(db),
		PhotoSizesDAO:     mysql_dao.NewPhotoSizesDAO(db),
		VideoSizesDAO:     mysql_dao.NewVideoSizesDAO(db),
		CommonDAO:         sqlx.NewCommonDAO(db),
	}
}

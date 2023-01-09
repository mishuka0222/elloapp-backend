package dao

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/media/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

func (m *Dao) SaveEncryptedFileV2(ctx context.Context, eF *mtproto.EncryptedFile) error {
	do := &dataobject.EncryptedFilesDO{
		EncryptedFileId: eF.Id,
		AccessHash:      eF.AccessHash,
		DcId:            eF.DcId,
		FilePath:        "",
		FileSize:        eF.GetFixedSize(),
		KeyFingerprint:  eF.KeyFingerprint,
		Md5Checksum:     "",
	}

	_, _, err := m.EncryptedFilesDAO.Insert(ctx, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("error - %v", err)
	}

	return err
}

func (m *Dao) GetEncryptedFile(ctx context.Context, id, accessHash int64) (*mtproto.EncryptedFile, error) {
	do, err := m.EncryptedFilesDAO.SelectByFileLocation(ctx, id, accessHash)
	if err != nil {
		logx.WithContext(ctx).Errorf("error - %v", err)
		return nil, err
	}

	if do == nil {
		return mtproto.MakeTLEncryptedFileEmpty(nil).To_EncryptedFile(), nil
	} else {
		encryptedFile := mtproto.MakeTLEncryptedFile(&mtproto.EncryptedFile{
			Id:             do.EncryptedFileId,
			AccessHash:     do.AccessHash,
			Size2_INT32:    int32(do.FileSize),
			Size2_INT64:    do.FileSize,
			DcId:           do.DcId,
			KeyFingerprint: do.KeyFingerprint,
		})
		return encryptedFile.To_EncryptedFile(), nil
	}
}

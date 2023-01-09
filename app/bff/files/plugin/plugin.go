package plugin

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type FilesPlugin interface {
	GetStickerSetThumbFileLocation(ctx context.Context, userId int64, stickerset *mtproto.InputStickerSet, version int32) (*mtproto.InputFileLocation, error)
	GetGroupCallStreamFile(ctx context.Context, userId int64, file *mtproto.InputFileLocation) (*mtproto.Upload_File, error)
}

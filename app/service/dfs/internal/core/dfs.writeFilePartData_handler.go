package core

import (
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/dfs"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/internal/model"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// DfsWriteFilePartData
// dfs.writeFilePartData flags:# creator:long file_id:long bytes:bytes big:flags.0?true file_total_parts:flags.1?int = Bool;
func (c *DfsCore) DfsWriteFilePartData(in *dfs.TLDfsWriteFilePartData) (*mtproto.Bool, error) {
	var (
		err error
	)

	if len(in.GetBytes()) == 0 {
		err = mtproto.ErrFilePartEmpty
		c.Logger.Errorf("dfs.writeFilePartData - error: %v", err)
		return nil, err
	}

	err = c.svcCtx.Dao.WriteFilePartData(c.ctx, in.Creator, in.FileId, in.FilePart, in.Bytes)
	if err != nil {
		c.Logger.Errorf("dfs.writeFilePartData - error: %v", err)
		return nil, err
	}

	if in.FilePart == 0 {
		err = c.svcCtx.Dao.SetFileInfo(c.ctx, &model.DfsFileInfo{
			Creator:           in.Creator,
			FileId:            in.FileId,
			Big:               in.Big,
			FileName:          "",
			FileTotalParts:    int(in.GetFileTotalParts().GetValue()),
			FirstFilePartSize: len(in.Bytes),
			FilePartSize:      0,
			LastFilePartSize:  0,
			MimeType:          "",
			Mtime:             time.Now().Unix(),
		})

		// TODO(@benqi): error
		if err != nil {
			c.Logger.Errorf("dfs.writeFilePartData - error: %v", err)
			return nil, mtproto.ErrFilePartInvalid
		}
	} else if in.FilePart == 1 {
		err = c.svcCtx.Dao.SetFileInfo(c.ctx, &model.DfsFileInfo{
			Creator:      in.Creator,
			FileId:       in.FileId,
			FilePartSize: len(in.Bytes),
		})

		// TODO(@benqi): error
		if err != nil {
			c.Logger.Errorf("dfs.writeFilePartData - error: %v", err)
			return nil, mtproto.ErrFilePartInvalid
		}
	}

	if in.GetFileTotalParts() != nil {
		if in.GetFileTotalParts().GetValue() == in.FilePart+1 {
			err = c.svcCtx.Dao.SetFileInfo(c.ctx, &model.DfsFileInfo{
				Creator:          in.Creator,
				FileId:           in.FileId,
				LastFilePartSize: len(in.Bytes),
			})

			// TODO(@benqi): error
			if err != nil {
				c.Logger.Errorf("dfs.writeFilePartData - error: %v", err)
				return nil, mtproto.ErrFilePartInvalid
			}
		}
	}

	return mtproto.BoolTrue, nil
}

package service

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/autodownload/internal/core"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AccountGetAutoDownloadSettings
// account.getAutoDownloadSettings#56da0b3f = account.AutoDownloadSettings;
func (s *Service) AccountGetAutoDownloadSettings(ctx context.Context, request *mtproto.TLAccountGetAutoDownloadSettings) (*mtproto.Account_AutoDownloadSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getAutoDownloadSettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountGetAutoDownloadSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getAutoDownloadSettings - reply: %s", r.DebugString())
	return r, err
}

// AccountSaveAutoDownloadSettings
// account.saveAutoDownloadSettings#76f36233 flags:# low:flags.0?true high:flags.1?true settings:AutoDownloadSettings = Bool;
func (s *Service) AccountSaveAutoDownloadSettings(ctx context.Context, request *mtproto.TLAccountSaveAutoDownloadSettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.saveAutoDownloadSettings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountSaveAutoDownloadSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.saveAutoDownloadSettings - reply: %s", r.DebugString())
	return r, err
}

package dao

import (
	"context"
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"

	"github.com/gogo/protobuf/types"
)

const (
	userPeerSettingsKeyPrefix = "user_peer_settings"
)

func genUserPeerSettingsCacheKey(id int64, peerType int32, peerId int64) string {
	return fmt.Sprintf("%s_%d_%d_%d", userPeerSettingsKeyPrefix, id, peerType, peerId)
}

func (d *Dao) GetUserPeerSettings(ctx context.Context, id int64, peerType int32, peerId int64) (*mtproto.PeerSettings, error) {
	settings := mtproto.MakeTLPeerSettings(nil).To_PeerSettings()

	err := d.CachedConn.QueryRow(
		ctx,
		settings,
		genUserPeerSettingsCacheKey(id, peerType, peerId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			settingsDO, err := d.UserPeerSettingsDAO.Select(ctx, id, peerType, peerId)
			if err != nil {
				return err
			}

			var (
				peerSettings = v.(*mtproto.PeerSettings)
			)

			if settingsDO != nil {
				peerSettings.ReportSpam = settingsDO.ReportSpam
				peerSettings.AddContact = settingsDO.AddContact
				peerSettings.BlockContact = settingsDO.BlockContact
				peerSettings.ShareContact = settingsDO.ShareContact
				peerSettings.NeedContactsException = settingsDO.NeedContactsException
				peerSettings.ReportGeo = settingsDO.ReportGeo
				peerSettings.Autoarchived = settingsDO.Autoarchived
				peerSettings.GeoDistance = nil

				if settingsDO.GeoDistance != 0 {
					peerSettings.GeoDistance = &types.Int32Value{Value: settingsDO.GeoDistance}
				}
			} else {
				peerSettings.ReportSpam = false
				peerSettings.AddContact = false
				peerSettings.BlockContact = false
				peerSettings.ShareContact = false
				peerSettings.NeedContactsException = false
				peerSettings.ReportGeo = false
				peerSettings.Autoarchived = false
				peerSettings.GeoDistance = nil
			}

			return nil
		})
	if err != nil {
		return nil, err
	}

	return settings, nil
}

func (d *Dao) SetUserPeerSettings(ctx context.Context, id int64, peerType int32, peerId int64, settings *mtproto.PeerSettings) error {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			return d.UserPeerSettingsDAO.InsertOrUpdate(
				ctx,
				&dataobject.UserPeerSettingsDO{
					UserId:                id,
					PeerType:              peerType,
					PeerId:                peerId,
					Hide:                  false,
					ReportSpam:            settings.ReportSpam,
					AddContact:            settings.AddContact,
					BlockContact:          settings.BlockContact,
					ShareContact:          settings.ShareContact,
					NeedContactsException: settings.NeedContactsException,
					ReportGeo:             settings.ReportGeo,
					Autoarchived:          settings.Autoarchived,
					GeoDistance:           settings.GetGeoDistance().GetValue(),
				})
		},
		genUserPeerSettingsCacheKey(id, peerType, peerId))

	return err
}

func (d *Dao) DeleteUserPeerSettings(ctx context.Context, id int64, peerType int32, peerId int64) error {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			affected, err := d.UserPeerSettingsDAO.Delete(ctx, id, peerType, peerId)
			return 0, affected, err
		},
		genUserPeerSettingsCacheKey(id, peerType, peerId))

	return err
}

package username_client

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

type UsernameClientHelper struct {
	cli UsernameClient
}

func NewUsernameClientHelper(cli zrpc.Client) *UsernameClientHelper {
	return &UsernameClientHelper{
		cli: NewUsernameClient(cli),
	}
}

func (m *UsernameClientHelper) Client() UsernameClient {
	return m.cli
}

func (m *UsernameClientHelper) CheckUsername(ctx context.Context, name string) (int, error) {
	rVal, err := m.cli.UsernameCheckUsername(ctx, &username.TLUsernameCheckUsername{
		Username: name,
	})
	if err != nil {
		return 0, err
	} else {
		switch rVal.GetPredicateName() {
		case username.Predicate_usernameNotExisted:
			return username.UsernameNotExisted, nil
		case username.Predicate_usernameExisted:
			return username.UsernameExisted, nil
		case username.Predicate_usernameExistedNotMe:
			return username.UsernameExistedNotMe, nil
		case username.Predicate_usernameExistedIsMe:
			return username.UsernameExistedIsMe, nil
		default:
			return username.UsernameNotExisted, nil
		}
	}
}

func (m *UsernameClientHelper) UpdateUsername(ctx context.Context, peerType int32, peerId int64, name string) (bool, error) {
	rB, err := m.cli.UsernameUpdateUsername(ctx, &username.TLUsernameUpdateUsername{
		PeerType: peerType,
		PeerId:   peerId,
		Username: name,
	})
	if err != nil {
		return false, err
	} else {
		return mtproto.FromBool(rB), nil
	}
}

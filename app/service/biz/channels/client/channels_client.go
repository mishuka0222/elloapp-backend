// Code generated by goctl. DO NOT EDIT.
// Source: channels.proto

package channels_client

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddChannelParticipantReq           = channels.AddChannelParticipantReq
	Channel                            = channels.Channel
	ChannelData                        = channels.ChannelData
	ChannelDataByIdReq                 = channels.ChannelDataByIdReq
	ChannelParticipant                 = channels.ChannelParticipant
	ChannelParticipantListReq          = channels.ChannelParticipantListReq
	ChannelParticipantListResp         = channels.ChannelParticipantListResp
	ChannelParticipantRes              = channels.ChannelParticipantRes
	ChannelParticipantsReq             = channels.ChannelParticipantsReq
	CheckChannelUserNameReq            = channels.CheckChannelUserNameReq
	CheckChannelUserNameResp           = channels.CheckChannelUserNameResp
	CheckDeleteChannelUserReq          = channels.CheckDeleteChannelUserReq
	CheckUserIsAdministratorReq        = channels.CheckUserIsAdministratorReq
	CheckUserIsAdministratorResp       = channels.CheckUserIsAdministratorResp
	CreateNewChannelReq                = channels.CreateNewChannelReq
	DeleteChannelUserReq               = channels.DeleteChannelUserReq
	EditChannelAdminReq                = channels.EditChannelAdminReq
	EditChannelPhotoReq                = channels.EditChannelPhotoReq
	EditChannelTitleReq                = channels.EditChannelTitleReq
	GetChannelBySelfIDReq              = channels.GetChannelBySelfIDReq
	GetChannelBySelfIDResp             = channels.GetChannelBySelfIDResp
	GetChannelFullBySelfIdReq          = channels.GetChannelFullBySelfIdReq
	GetChannelFullBySelfIdResp         = channels.GetChannelFullBySelfIdResp
	GetChannelParticipantIdListDaoReq  = channels.GetChannelParticipantIdListDaoReq
	GetChannelParticipantIdListDaoResp = channels.GetChannelParticipantIdListDaoResp
	GetChannelParticipantIdListResp    = channels.GetChannelParticipantIdListResp
	GetChannelParticipantsResp         = channels.GetChannelParticipantsResp
	GetChatsListBySelfAndIDListReq     = channels.GetChatsListBySelfAndIDListReq
	GetChatsListBySelfAndIDListResp    = channels.GetChatsListBySelfAndIDListResp
	MakeChannelParticipant2ByDOReq     = channels.MakeChannelParticipant2ByDOReq
	ToChatReq                          = channels.ToChatReq
	ToChatResp                         = channels.ToChatResp
	ToggleChannelAdminsReq             = channels.ToggleChannelAdminsReq
	UpdateChannelLinkReq               = channels.UpdateChannelLinkReq
	UpdateChannelLinkResp              = channels.UpdateChannelLinkResp
	Void                               = channels.Void

	ChannelsClient interface {
		// NEW
		GetChannelDataById(ctx context.Context, in *ChannelDataByIdReq, opts ...grpc.CallOption) (*ChannelData, error)
		CreateNewChannel(ctx context.Context, in *CreateNewChannelReq, opts ...grpc.CallOption) (*ChannelData, error)
		AddChannelParticipant(ctx context.Context, in *AddChannelParticipantReq, opts ...grpc.CallOption) (*Void, error)
		GetChatsListBySelfAndIDList(ctx context.Context, in *GetChatsListBySelfAndIDListReq, opts ...grpc.CallOption) (*GetChatsListBySelfAndIDListResp, error)
		GetChannelFullBySelfId(ctx context.Context, in *GetChannelFullBySelfIdReq, opts ...grpc.CallOption) (*GetChannelFullBySelfIdResp, error)
		GetChannelParticipantList(ctx context.Context, in *ChannelParticipantListReq, opts ...grpc.CallOption) (*ChannelParticipantListResp, error)
		GetChannelParticipants(ctx context.Context, in *ChannelParticipantsReq, opts ...grpc.CallOption) (*GetChannelParticipantsResp, error)
		ToChat(ctx context.Context, in *ToChatReq, opts ...grpc.CallOption) (*ToChatResp, error)
		CheckChannelUserName(ctx context.Context, in *CheckChannelUserNameReq, opts ...grpc.CallOption) (*CheckChannelUserNameResp, error)
		UpdateChannelLink(ctx context.Context, in *UpdateChannelLinkReq, opts ...grpc.CallOption) (*UpdateChannelLinkResp, error)
		CheckUserIsAdministrator(ctx context.Context, in *CheckUserIsAdministratorReq, opts ...grpc.CallOption) (*CheckUserIsAdministratorResp, error)
		EditChannelAdmin(ctx context.Context, in *EditChannelAdminReq, opts ...grpc.CallOption) (*Void, error)
		// In Work
		DeleteChannelUser(ctx context.Context, in *DeleteChannelUserReq, opts ...grpc.CallOption) (*Void, error)
		EditChannelTitle(ctx context.Context, in *EditChannelTitleReq, opts ...grpc.CallOption) (*Void, error)
		EditChannelPhoto(ctx context.Context, in *EditChannelPhotoReq, opts ...grpc.CallOption) (*Void, error)
		// OLD
		ToggleChannelAdmins(ctx context.Context, in *ToggleChannelAdminsReq, opts ...grpc.CallOption) (*Void, error)
		// channel_data
		GetChannelParticipantIdList(ctx context.Context, in *ChannelData, opts ...grpc.CallOption) (*GetChannelParticipantIdListResp, error)
		GetChannelBySelfID(ctx context.Context, in *GetChannelBySelfIDReq, opts ...grpc.CallOption) (*GetChannelBySelfIDResp, error)
		GetChannelParticipantIdListDao(ctx context.Context, in *GetChannelParticipantIdListDaoReq, opts ...grpc.CallOption) (*GetChannelParticipantIdListDaoResp, error)
	}

	defaultChannelsClient struct {
		cli zrpc.Client
	}
)

func NewChannelsClient(cli zrpc.Client) ChannelsClient {
	return &defaultChannelsClient{
		cli: cli,
	}
}

// NEW
func (m *defaultChannelsClient) GetChannelDataById(ctx context.Context, in *ChannelDataByIdReq, opts ...grpc.CallOption) (*ChannelData, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelDataById(ctx, in, opts...)
}

func (m *defaultChannelsClient) CreateNewChannel(ctx context.Context, in *CreateNewChannelReq, opts ...grpc.CallOption) (*ChannelData, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.CreateNewChannel(ctx, in, opts...)
}

func (m *defaultChannelsClient) AddChannelParticipant(ctx context.Context, in *AddChannelParticipantReq, opts ...grpc.CallOption) (*Void, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.AddChannelParticipant(ctx, in, opts...)
}

func (m *defaultChannelsClient) GetChatsListBySelfAndIDList(ctx context.Context, in *GetChatsListBySelfAndIDListReq, opts ...grpc.CallOption) (*GetChatsListBySelfAndIDListResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChatsListBySelfAndIDList(ctx, in, opts...)
}

func (m *defaultChannelsClient) GetChannelFullBySelfId(ctx context.Context, in *GetChannelFullBySelfIdReq, opts ...grpc.CallOption) (*GetChannelFullBySelfIdResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelFullBySelfId(ctx, in, opts...)
}

func (m *defaultChannelsClient) GetChannelParticipantList(ctx context.Context, in *ChannelParticipantListReq, opts ...grpc.CallOption) (*ChannelParticipantListResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelParticipantList(ctx, in, opts...)
}

func (m *defaultChannelsClient) GetChannelParticipants(ctx context.Context, in *ChannelParticipantsReq, opts ...grpc.CallOption) (*GetChannelParticipantsResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelParticipants(ctx, in, opts...)
}

func (m *defaultChannelsClient) ToChat(ctx context.Context, in *ToChatReq, opts ...grpc.CallOption) (*ToChatResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.ToChat(ctx, in, opts...)
}

func (m *defaultChannelsClient) CheckChannelUserName(ctx context.Context, in *CheckChannelUserNameReq, opts ...grpc.CallOption) (*CheckChannelUserNameResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.CheckChannelUserName(ctx, in, opts...)
}

func (m *defaultChannelsClient) UpdateChannelLink(ctx context.Context, in *UpdateChannelLinkReq, opts ...grpc.CallOption) (*UpdateChannelLinkResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.UpdateChannelLink(ctx, in, opts...)
}

func (m *defaultChannelsClient) CheckUserIsAdministrator(ctx context.Context, in *CheckUserIsAdministratorReq, opts ...grpc.CallOption) (*CheckUserIsAdministratorResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.CheckUserIsAdministrator(ctx, in, opts...)
}

func (m *defaultChannelsClient) EditChannelAdmin(ctx context.Context, in *EditChannelAdminReq, opts ...grpc.CallOption) (*Void, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.EditChannelAdmin(ctx, in, opts...)
}

// In Work
func (m *defaultChannelsClient) DeleteChannelUser(ctx context.Context, in *DeleteChannelUserReq, opts ...grpc.CallOption) (*Void, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.DeleteChannelUser(ctx, in, opts...)
}

func (m *defaultChannelsClient) EditChannelTitle(ctx context.Context, in *EditChannelTitleReq, opts ...grpc.CallOption) (*Void, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.EditChannelTitle(ctx, in, opts...)
}

func (m *defaultChannelsClient) EditChannelPhoto(ctx context.Context, in *EditChannelPhotoReq, opts ...grpc.CallOption) (*Void, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.EditChannelPhoto(ctx, in, opts...)
}

// OLD
func (m *defaultChannelsClient) ToggleChannelAdmins(ctx context.Context, in *ToggleChannelAdminsReq, opts ...grpc.CallOption) (*Void, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.ToggleChannelAdmins(ctx, in, opts...)
}

// channel_data
func (m *defaultChannelsClient) GetChannelParticipantIdList(ctx context.Context, in *ChannelData, opts ...grpc.CallOption) (*GetChannelParticipantIdListResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelParticipantIdList(ctx, in, opts...)
}

func (m *defaultChannelsClient) GetChannelBySelfID(ctx context.Context, in *GetChannelBySelfIDReq, opts ...grpc.CallOption) (*GetChannelBySelfIDResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelBySelfID(ctx, in, opts...)
}

func (m *defaultChannelsClient) GetChannelParticipantIdListDao(ctx context.Context, in *GetChannelParticipantIdListDaoReq, opts ...grpc.CallOption) (*GetChannelParticipantIdListDaoResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelParticipantIdListDao(ctx, in, opts...)
}

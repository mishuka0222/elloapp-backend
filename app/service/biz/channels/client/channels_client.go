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
	AddChannelUserReq                  = channels.AddChannelUserReq
	AddChannelUserResp                 = channels.AddChannelUserResp
	Channel                            = channels.Channel
	ChannelCoreByCreateChannelReq      = channels.ChannelCoreByCreateChannelReq
	ChannelCoreByIdReq                 = channels.ChannelCoreByIdReq
	ChannelCoreData                    = channels.ChannelCoreData
	ChannelMessageBox                  = channels.ChannelMessageBox
	ChannelParticipant                 = channels.ChannelParticipant
	ChannelParticipantRes              = channels.ChannelParticipantRes
	CheckChannelUserNameReq            = channels.CheckChannelUserNameReq
	CheckChannelUserNameResp           = channels.CheckChannelUserNameResp
	CheckDeleteChannelUserReq          = channels.CheckDeleteChannelUserReq
	CheckDeleteChannelUserResp         = channels.CheckDeleteChannelUserResp
	CheckUserIsAdministratorReq        = channels.CheckUserIsAdministratorReq
	CheckUserIsAdministratorResp       = channels.CheckUserIsAdministratorResp
	CreateChannelMessageBoxReq         = channels.CreateChannelMessageBoxReq
	DeleteChannelUserReq               = channels.DeleteChannelUserReq
	DeleteChannelUserResp              = channels.DeleteChannelUserResp
	DoToChannelMessageRes              = channels.DoToChannelMessageRes
	EditChannelAdminReq                = channels.EditChannelAdminReq
	EditChannelAdminResp               = channels.EditChannelAdminResp
	EditChannelPhotoReq                = channels.EditChannelPhotoReq
	EditChannelPhotoResp               = channels.EditChannelPhotoResp
	EditChannelTitleReq                = channels.EditChannelTitleReq
	EditChannelTitleResp               = channels.EditChannelTitleResp
	ExportedChatInviteResp             = channels.ExportedChatInviteResp
	GetChannelBySelfIDReq              = channels.GetChannelBySelfIDReq
	GetChannelBySelfIDResp             = channels.GetChannelBySelfIDResp
	GetChannelFullBySelfIdReq          = channels.GetChannelFullBySelfIdReq
	GetChannelFullBySelfIdResp         = channels.GetChannelFullBySelfIdResp
	GetChannelListBySelfAndIDListReq   = channels.GetChannelListBySelfAndIDListReq
	GetChannelListBySelfAndIDListResp  = channels.GetChannelListBySelfAndIDListResp
	GetChannelMessageReq               = channels.GetChannelMessageReq
	GetChannelMessageRes               = channels.GetChannelMessageRes
	GetChannelParticipantIdListDaoReq  = channels.GetChannelParticipantIdListDaoReq
	GetChannelParticipantIdListDaoResp = channels.GetChannelParticipantIdListDaoResp
	GetChannelParticipantIdListResp    = channels.GetChannelParticipantIdListResp
	GetChannelParticipantListResp      = channels.GetChannelParticipantListResp
	GetChannelParticipantsResp         = channels.GetChannelParticipantsResp
	MakeChannelParticipant2ByDOReq     = channels.MakeChannelParticipant2ByDOReq
	MessageData                        = channels.MessageData
	ToChannelReq                       = channels.ToChannelReq
	ToChannelResp                      = channels.ToChannelResp
	ToggleChannelAdminsReq             = channels.ToggleChannelAdminsReq
	ToggleChannelAdminsResp            = channels.ToggleChannelAdminsResp

	ChannelsClient interface {
		// channel_data
		MakeChannelParticipant2ByDO(ctx context.Context, in *MakeChannelParticipant2ByDOReq, opts ...grpc.CallOption) (*ChannelParticipantRes, error)
		NewChannelCoreById(ctx context.Context, in *ChannelCoreByIdReq, opts ...grpc.CallOption) (*ChannelCoreData, error)
		NewChannelCoreByCreateChannel(ctx context.Context, in *ChannelCoreByCreateChannelReq, opts ...grpc.CallOption) (*ChannelCoreData, error)
		ExportedChatInvite(ctx context.Context, in *ChannelCoreData, opts ...grpc.CallOption) (*ExportedChatInviteResp, error)
		CheckUserIsAdministrator(ctx context.Context, in *CheckUserIsAdministratorReq, opts ...grpc.CallOption) (*CheckUserIsAdministratorResp, error)
		GetChannelParticipantList(ctx context.Context, in *ChannelCoreData, opts ...grpc.CallOption) (*GetChannelParticipantListResp, error)
		GetChannelParticipantIdList(ctx context.Context, in *ChannelCoreData, opts ...grpc.CallOption) (*GetChannelParticipantIdListResp, error)
		GetChannelParticipants(ctx context.Context, in *ChannelCoreData, opts ...grpc.CallOption) (*GetChannelParticipantsResp, error)
		AddChannelUser(ctx context.Context, in *AddChannelUserReq, opts ...grpc.CallOption) (*AddChannelUserResp, error)
		ToChannel(ctx context.Context, in *ToChannelReq, opts ...grpc.CallOption) (*ToChannelResp, error)
		CheckDeleteChannelUser(ctx context.Context, in *CheckDeleteChannelUserReq, opts ...grpc.CallOption) (*CheckDeleteChannelUserResp, error)
		DeleteChannelUser(ctx context.Context, in *DeleteChannelUserReq, opts ...grpc.CallOption) (*DeleteChannelUserResp, error)
		EditChannelTitle(ctx context.Context, in *EditChannelTitleReq, opts ...grpc.CallOption) (*EditChannelTitleResp, error)
		EditChannelPhoto(ctx context.Context, in *EditChannelPhotoReq, opts ...grpc.CallOption) (*EditChannelPhotoResp, error)
		EditChannelAdmin(ctx context.Context, in *EditChannelAdminReq, opts ...grpc.CallOption) (*EditChannelAdminResp, error)
		ToggleChannelAdmins(ctx context.Context, in *ToggleChannelAdminsReq, opts ...grpc.CallOption) (*ToggleChannelAdminsResp, error)
		// channel_util
		GetChannelListBySelfAndIDList(ctx context.Context, in *GetChannelListBySelfAndIDListReq, opts ...grpc.CallOption) (*GetChannelListBySelfAndIDListResp, error)
		CheckChannelUserName(ctx context.Context, in *CheckChannelUserNameReq, opts ...grpc.CallOption) (*CheckChannelUserNameResp, error)
		GetChannelBySelfID(ctx context.Context, in *GetChannelBySelfIDReq, opts ...grpc.CallOption) (*GetChannelBySelfIDResp, error)
		GetChannelParticipantIdListDao(ctx context.Context, in *GetChannelParticipantIdListDaoReq, opts ...grpc.CallOption) (*GetChannelParticipantIdListDaoResp, error)
		GetChannelFullBySelfId(ctx context.Context, in *GetChannelFullBySelfIdReq, opts ...grpc.CallOption) (*GetChannelFullBySelfIdResp, error)
		// channel_message_box
		CreateChannelMessageBox(ctx context.Context, in *CreateChannelMessageBoxReq, opts ...grpc.CallOption) (*ChannelMessageBox, error)
		DoToChannelMessage(ctx context.Context, in *MessageData, opts ...grpc.CallOption) (*DoToChannelMessageRes, error)
		GetChannelMessage(ctx context.Context, in *GetChannelMessageReq, opts ...grpc.CallOption) (*GetChannelMessageRes, error)
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

// channel_data
func (m *defaultChannelsClient) MakeChannelParticipant2ByDO(ctx context.Context, in *MakeChannelParticipant2ByDOReq, opts ...grpc.CallOption) (*ChannelParticipantRes, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.MakeChannelParticipant2ByDO(ctx, in, opts...)
}

func (m *defaultChannelsClient) NewChannelCoreById(ctx context.Context, in *ChannelCoreByIdReq, opts ...grpc.CallOption) (*ChannelCoreData, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.NewChannelCoreById(ctx, in, opts...)
}

func (m *defaultChannelsClient) NewChannelCoreByCreateChannel(ctx context.Context, in *ChannelCoreByCreateChannelReq, opts ...grpc.CallOption) (*ChannelCoreData, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.NewChannelCoreByCreateChannel(ctx, in, opts...)
}

func (m *defaultChannelsClient) ExportedChatInvite(ctx context.Context, in *ChannelCoreData, opts ...grpc.CallOption) (*ExportedChatInviteResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.ExportedChatInvite(ctx, in, opts...)
}

func (m *defaultChannelsClient) CheckUserIsAdministrator(ctx context.Context, in *CheckUserIsAdministratorReq, opts ...grpc.CallOption) (*CheckUserIsAdministratorResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.CheckUserIsAdministrator(ctx, in, opts...)
}

func (m *defaultChannelsClient) GetChannelParticipantList(ctx context.Context, in *ChannelCoreData, opts ...grpc.CallOption) (*GetChannelParticipantListResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelParticipantList(ctx, in, opts...)
}

func (m *defaultChannelsClient) GetChannelParticipantIdList(ctx context.Context, in *ChannelCoreData, opts ...grpc.CallOption) (*GetChannelParticipantIdListResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelParticipantIdList(ctx, in, opts...)
}

func (m *defaultChannelsClient) GetChannelParticipants(ctx context.Context, in *ChannelCoreData, opts ...grpc.CallOption) (*GetChannelParticipantsResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelParticipants(ctx, in, opts...)
}

func (m *defaultChannelsClient) AddChannelUser(ctx context.Context, in *AddChannelUserReq, opts ...grpc.CallOption) (*AddChannelUserResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.AddChannelUser(ctx, in, opts...)
}

func (m *defaultChannelsClient) ToChannel(ctx context.Context, in *ToChannelReq, opts ...grpc.CallOption) (*ToChannelResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.ToChannel(ctx, in, opts...)
}

func (m *defaultChannelsClient) CheckDeleteChannelUser(ctx context.Context, in *CheckDeleteChannelUserReq, opts ...grpc.CallOption) (*CheckDeleteChannelUserResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.CheckDeleteChannelUser(ctx, in, opts...)
}

func (m *defaultChannelsClient) DeleteChannelUser(ctx context.Context, in *DeleteChannelUserReq, opts ...grpc.CallOption) (*DeleteChannelUserResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.DeleteChannelUser(ctx, in, opts...)
}

func (m *defaultChannelsClient) EditChannelTitle(ctx context.Context, in *EditChannelTitleReq, opts ...grpc.CallOption) (*EditChannelTitleResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.EditChannelTitle(ctx, in, opts...)
}

func (m *defaultChannelsClient) EditChannelPhoto(ctx context.Context, in *EditChannelPhotoReq, opts ...grpc.CallOption) (*EditChannelPhotoResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.EditChannelPhoto(ctx, in, opts...)
}

func (m *defaultChannelsClient) EditChannelAdmin(ctx context.Context, in *EditChannelAdminReq, opts ...grpc.CallOption) (*EditChannelAdminResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.EditChannelAdmin(ctx, in, opts...)
}

func (m *defaultChannelsClient) ToggleChannelAdmins(ctx context.Context, in *ToggleChannelAdminsReq, opts ...grpc.CallOption) (*ToggleChannelAdminsResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.ToggleChannelAdmins(ctx, in, opts...)
}

// channel_util
func (m *defaultChannelsClient) GetChannelListBySelfAndIDList(ctx context.Context, in *GetChannelListBySelfAndIDListReq, opts ...grpc.CallOption) (*GetChannelListBySelfAndIDListResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelListBySelfAndIDList(ctx, in, opts...)
}

func (m *defaultChannelsClient) CheckChannelUserName(ctx context.Context, in *CheckChannelUserNameReq, opts ...grpc.CallOption) (*CheckChannelUserNameResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.CheckChannelUserName(ctx, in, opts...)
}

func (m *defaultChannelsClient) GetChannelBySelfID(ctx context.Context, in *GetChannelBySelfIDReq, opts ...grpc.CallOption) (*GetChannelBySelfIDResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelBySelfID(ctx, in, opts...)
}

func (m *defaultChannelsClient) GetChannelParticipantIdListDao(ctx context.Context, in *GetChannelParticipantIdListDaoReq, opts ...grpc.CallOption) (*GetChannelParticipantIdListDaoResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelParticipantIdListDao(ctx, in, opts...)
}

func (m *defaultChannelsClient) GetChannelFullBySelfId(ctx context.Context, in *GetChannelFullBySelfIdReq, opts ...grpc.CallOption) (*GetChannelFullBySelfIdResp, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelFullBySelfId(ctx, in, opts...)
}

// channel_message_box
func (m *defaultChannelsClient) CreateChannelMessageBox(ctx context.Context, in *CreateChannelMessageBoxReq, opts ...grpc.CallOption) (*ChannelMessageBox, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.CreateChannelMessageBox(ctx, in, opts...)
}

func (m *defaultChannelsClient) DoToChannelMessage(ctx context.Context, in *MessageData, opts ...grpc.CallOption) (*DoToChannelMessageRes, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.DoToChannelMessage(ctx, in, opts...)
}

func (m *defaultChannelsClient) GetChannelMessage(ctx context.Context, in *GetChannelMessageReq, opts ...grpc.CallOption) (*GetChannelMessageRes, error) {
	client := channels.NewRPCChannelsClient(m.cli.Conn())
	return client.GetChannelMessage(ctx, in, opts...)
}

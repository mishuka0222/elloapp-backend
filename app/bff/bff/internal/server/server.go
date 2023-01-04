package server

import (
	"flag"
	account_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/account"
	authorization_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization"
	autodownload_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/autodownload"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/bff/internal/config"
	bizraw_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/bizraw"
	op_srv "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/bizraw/service"
	chatinvites_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/chatinvites"
	chats_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/chats"
	configuration_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/configuration"
	contacts_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/contacts"
	dialogs_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/dialogs"
	drafts_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/drafts"
	feeds_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/feeds"
	files_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/files"
	messages_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/messages"
	miscellaneous_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/miscellaneous"
	notification_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/notification"
	nsfw_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/nsfw"
	photos_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/photos"
	premium_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/premium"
	qrcode_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/qrcode"
	secretchats_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/secretchats"
	sponsoredmessages_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/sponsoredmessages"
	tos_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/tos"
	updates_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/updates"
	usernames_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/usernames"
	users_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/users"
	voipcalls_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/voipcalls"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/bff.yaml", "the config file")

type Server struct {
	grpcSrv *zrpc.RpcServer
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.Infov(c)
	// ctx := svc.NewServiceContext(c)
	// s.grpcSrv = grpc.New(ctx, c.RpcServerConf)

	s.grpcSrv = zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {

		// secretchats
		mtproto.RegisterRPCSecretChatsServer(
			grpcServer,
			secretchats_helper.New(
				secretchats_helper.Config{}))

		// voipcalls
		mtproto.RegisterRPCVoipCallsServer(
			grpcServer,
			voipcalls_helper.New(
				voipcalls_helper.Config{}))

		// tos_helper
		mtproto.RegisterRPCTosServer(
			grpcServer,
			tos_helper.New(tos_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// configuration_helper
		mtproto.RegisterRPCConfigurationServer(
			grpcServer,
			configuration_helper.New(configuration_helper.Config{
				RpcServerConf:       c.RpcServerConf,
				ConfigurationClient: c.BizServiceClient,
			}))

		// qrcode_helper
		mtproto.RegisterRPCQrCodeServer(
			grpcServer,
			qrcode_helper.New(qrcode_helper.Config{
				RpcServerConf:     c.RpcServerConf,
				KV:                c.KV,
				UserClient:        c.BizServiceClient,
				AuthSessionClient: c.AuthSessionClient,
				SyncClient:        c.SyncClient,
			}))

		// miscellaneous_helper
		mtproto.RegisterRPCMiscellaneousServer(
			grpcServer,
			miscellaneous_helper.New(miscellaneous_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// authorization_helper
		authorizationService := authorization_helper.New(
			authorization_helper.Config{
				RpcServerConf:     c.RpcServerConf,
				KV:                c.KV,
				Code:              c.Code,
				UserClient:        c.BizServiceClient,
				UsernameClient:    c.BizServiceClient,
				AuthsessionClient: c.AuthSessionClient,
				ChatClient:        c.BizServiceClient,
				StatusClient:      c.StatusClient,
				SyncClient:        c.SyncClient,
				MsgClient:         c.MsgClient,
			},
			nil,
			nil)
		mtproto.RegisterRPCAuthorizationServer(
			grpcServer,
			authorizationService)

		//authorizationCustom := authorization_customize_helper.New(
		//	authorization_customize_helper.Config{
		//		AuthorizationClient: c.BizServiceClient,
		//	}, authorizationService)

		// premium_helper
		mtproto.RegisterRPCPremiumServer(
			grpcServer,
			premium_helper.New(premium_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// chatinvites_helper
		mtproto.RegisterRPCChatInvitesServer(
			grpcServer,
			chatinvites_helper.New(chatinvites_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UserClient:    c.BizServiceClient,
				ChatClient:    c.BizServiceClient,
				MsgClient:     c.MsgClient,
			}))

		// chats_helper
		mtproto.RegisterRPCChatsServer(
			grpcServer,
			chats_helper.New(chats_helper.Config{
				RpcServerConf:     c.RpcServerConf,
				UserClient:        c.BizServiceClient,
				ChatClient:        c.BizServiceClient,
				MsgClient:         c.MsgClient,
				DialogClient:      c.BizServiceClient,
				SyncClient:        c.SyncClient,
				MediaClient:       c.MediaClient,
				AuthsessionClient: c.AuthSessionClient,
				IdgenClient:       c.IdgenClient,
				MessageClient:     c.BizServiceClient,
			}))

		// files_helper
		mtproto.RegisterRPCFilesServer(
			grpcServer,
			files_helper.New(files_helper.Config{
				RpcServerConf: c.RpcServerConf,
				DfsClient:     c.DfsClient,
				UserClient:    c.BizServiceClient,
				MediaClient:   c.MediaClient,
			}, nil))

		// updates_helper
		mtproto.RegisterRPCUpdatesServer(
			grpcServer,
			updates_helper.New(updates_helper.Config{
				RpcServerConf:     c.RpcServerConf,
				UpdatesClient:     c.BizServiceClient,
				UserClient:        c.BizServiceClient,
				ChatClient:        c.BizServiceClient,
				AuthsessionClient: c.AuthSessionClient,
			}))

		// contacts_helper
		mtproto.RegisterRPCContactsServer(
			grpcServer,
			contacts_helper.New(
				contacts_helper.Config{
					RpcServerConf:  c.RpcServerConf,
					UserClient:     c.BizServiceClient,
					ChatClient:     c.BizServiceClient,
					UsernameClient: c.BizServiceClient,
					SyncClient:     c.SyncClient,
				},
				nil))

		// dialogs_helper
		mtproto.RegisterRPCDialogsServer(
			grpcServer,
			dialogs_helper.New(dialogs_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UpdatesClient: c.BizServiceClient,
				UserClient:    c.BizServiceClient,
				ChatClient:    c.BizServiceClient,
				DialogClient:  c.BizServiceClient,
				SyncClient:    c.SyncClient,
				MessageClient: c.BizServiceClient,
			}, nil))

		// drafts_helper
		mtproto.RegisterRPCDraftsServer(
			grpcServer,
			drafts_helper.New(drafts_helper.Config{
				RpcServerConf: c.RpcServerConf,
				DialogClient:  c.BizServiceClient,
				UserClient:    c.BizServiceClient,
				SyncClient:    c.SyncClient,
				ChatClient:    c.BizServiceClient,
			}, nil))

		// autodownload_helper
		mtproto.RegisterRPCAutoDownloadServer(
			grpcServer,
			autodownload_helper.New(autodownload_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// messages_helper

		messagesCore := messages_helper.New(messages_helper.Config{
			RpcServerConf:  c.RpcServerConf,
			UserClient:     c.BizServiceClient,
			ChatClient:     c.BizServiceClient,
			MsgClient:      c.MsgClient,
			DialogClient:   c.BizServiceClient,
			IdgenClient:    c.IdgenClient,
			MessageClient:  c.BizServiceClient,
			MediaClient:    c.MediaClient,
			UsernameClient: c.BizServiceClient,
			SyncClient:     c.SyncClient,
		}, nil)
		mtproto.RegisterRPCMessagesServer(
			grpcServer, messagesCore,
		)

		// notification_helper
		mtproto.RegisterRPCNotificationServer(
			grpcServer,
			notification_helper.New(notification_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UserClient:    c.BizServiceClient,
				ChatClient:    c.BizServiceClient,
				SyncClient:    c.SyncClient,
			}, nil))

		// users_helper
		mtproto.RegisterRPCUsersServer(
			grpcServer,
			users_helper.New(users_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UserClient:    c.BizServiceClient,
				ChatClient:    c.BizServiceClient,
				DialogClient:  c.BizServiceClient,
			}))

		// nsfw_helper
		mtproto.RegisterRPCNsfwServer(
			grpcServer,
			nsfw_helper.New(nsfw_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UserClient:    c.BizServiceClient,
			}))

		// sponsoredmessages_helper
		mtproto.RegisterRPCSponsoredMessagesServer(
			grpcServer,
			sponsoredmessages_helper.New(sponsoredmessages_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// account_helper
		mtproto.RegisterRPCAccountServer(
			grpcServer,
			account_helper.New(account_helper.Config{
				RpcServerConf:     c.RpcServerConf,
				UserClient:        c.BizServiceClient,
				AuthsessionClient: c.AuthSessionClient,
				ChatClient:        c.BizServiceClient,
				SyncClient:        c.SyncClient,
			}))

		// photos_helper
		mtproto.RegisterRPCPhotosServer(
			grpcServer,
			photos_helper.New(photos_helper.Config{
				RpcServerConf: c.RpcServerConf,
				MediaClient:   c.MediaClient,
				UserClient:    c.BizServiceClient,
				SyncClient:    c.SyncClient,
			}))

		// usernames_helper
		mtproto.RegisterRPCUsernamesServer(
			grpcServer,
			usernames_helper.New(usernames_helper.Config{
				RpcServerConf:  c.RpcServerConf,
				UserClient:     c.BizServiceClient,
				UsernameClient: c.BizServiceClient,
				ChatClient:     c.BizServiceClient,
				SyncClient:     c.SyncClient,
			}, nil))

		// bizraw_helper
		mtproto.RegisterRPCBizServer(
			grpcServer,
			bizraw_helper.New(
				bizraw_helper.Config{
					RpcServerConf: c.RpcServerConf,
				}, map[op_srv.ServiceID]op_srv.OperationServer{
					op_srv.Feeds: feeds_helper.New(feeds_helper.Config{
						MessageClient: c.BizServiceClient,
						FeedsClient:   c.BizServiceClient,
					}, messagesCore),
					//op_srv.AuthorizationCustomize: authorizationCustom,
				}))

	})

	// logx.Must(err)

	go func() {
		s.grpcSrv.Start()
	}()
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
}

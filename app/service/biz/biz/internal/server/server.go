package server

import (
	"flag"

	accounts_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/accounts"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/accounts/accounts"
	authorization_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/biz/internal/config"
	channels_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	chat_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	code_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/code"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/code/code"
	configuration_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/configuration"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/configuration/configuration"
	dialog_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	feeds_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/feeds"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/feeds/feeds"
	message_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/message"
	phonecall_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/phonecall"
	updates_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/updates"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/updates/updates"
	user_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	username_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/username"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/biz.yaml", "the config file")

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

		//phone_call
		phonecall.RegisterRPCPhoneCallServer(
			grpcServer,
			phonecall_helper.New(
				phonecall_helper.Config{
					RpcServerConf: c.RpcServerConf,
					IdgenClient:   c.IdgenClient,
				}))

		//configuration
		configuration.RegisterRPCConfigurationServer(
			grpcServer,
			configuration_helper.New(
				configuration_helper.Config{
					RpcServerConf: c.RpcServerConf,
					Mysql:         c.Mysql,
				}))

		// feeds
		feeds.RegisterRPCFeedsServer(
			grpcServer,
			feeds_helper.New(
				feeds_helper.Config{
					RpcServerConf: c.RpcServerConf,
					Mysql:         c.Mysql,
				}))

		// accounts
		accounts.RegisterRPCAccountsServer(
			grpcServer,
			accounts_helper.New(
				accounts_helper.Config{
					RpcServerConf: c.RpcServerConf,
					Mysql:         c.Mysql,
				}))

		// chat_helper
		chat.RegisterRPCChatServer(
			grpcServer,
			chat_helper.New(
				chat_helper.Config{
					RpcServerConf: c.RpcServerConf,
					Mysql:         c.Mysql,
					Cache:         c.Cache,
					MediaClient:   c.MediaClient,
				},
				nil))

		// code_helper
		code.RegisterRPCCodeServer(
			grpcServer,
			code_helper.New(code_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
				KV:            c.KV,
			}))

		// dialog_helper
		dialog.RegisterRPCDialogServer(
			grpcServer,
			dialog_helper.New(dialog_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				Cache:         c.Cache,
			}))

		// message_helper
		message.RegisterRPCMessageServer(
			grpcServer,
			message_helper.New(
				message_helper.Config{
					RpcServerConf:   c.RpcServerConf,
					Mysql:           c.Mysql,
					Cache:           c.Cache,
					MessageSharding: c.MessageSharding,
				},
				nil))

		// updates_helper
		updates.RegisterRPCUpdatesServer(
			grpcServer,
			updates_helper.New(updates_helper.Config{
				RpcServerConf: c.RpcServerConf,
				Mysql:         c.Mysql,
				KV:            c.KV,
				IdgenClient:   c.IdgenClient,
			}))

		userService := user_helper.New(user_helper.Config{
			RpcServerConf: c.RpcServerConf,
			Mysql:         c.Mysql,
			Cache:         c.Cache,
			MediaClient:   c.MediaClient,
		})
		// user_helper
		user.RegisterRPCUserServer(
			grpcServer,
			userService,
		)

		// authorization
		authorization.RegisterRPCAuthorizationServer(
			grpcServer,
			authorization_helper.New(
				authorization_helper.Config{
					RpcServerConf:     c.RpcServerConf,
					Mysql:             c.Mysql,
					AuthSessionClient: c.AuthSessionClient,
					SyncClient:        c.SyncClient,
				}, userService))

		// username_helper
		usernameService := username_helper.New(username_helper.Config{
			RpcServerConf: c.RpcServerConf,
			Mysql:         c.Mysql,
			Cache:         c.Cache,
		})
		username.RegisterRPCUsernameServer(
			grpcServer,
			usernameService,
		)

		// channels
		channels.RegisterRPCChannelsServer(
			grpcServer,
			channels_helper.New(
				channels_helper.Config{
					RpcServerConf: c.RpcServerConf,
					Mysql:         c.Mysql,
					IdgenClient:   c.IdgenClient,
					MediaClient:   c.MediaClient,
				}, userService, usernameService))
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

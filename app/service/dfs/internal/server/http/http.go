package http

import (
	"github.com/zeromicro/go-zero/rest"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/svc"
	"net/http"
)

// New new a grpc server.
func New(ctx *svc.ServiceContext, c rest.RestConf) *rest.Server {
	srv := rest.MustNewServer(c)

	go func() {
		defer srv.Stop()

		srv.AddRoutes([]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/dfs/file/:file",
				Handler: getDfsFile(ctx),
			},
		})

		srv.Start()
	}()
	return srv
}

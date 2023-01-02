package main

import (
	"context"
	"io/ioutil"
	"os"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/internal/model"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/commands"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	c config.Config
)

func main() {
	commands.Run(New())
}

type Server struct {
	dao *dao.Dao
}

func New() *Server {
	s := new(Server)
	return s
}

func (s *Server) Initialize() error {
	s.dao = dao.New(c)
	return nil
}

func (s *Server) RunLoop() {
	defer func() {
		s.Destroy()
		os.Exit(0)
	}()

	logx.Infof("aaaa", "aa")
	// buf, err := ioutil.ReadFile("./test001.jpeg")
	// ../../../../../../../tools/gif2mp4/safe_image.gif
	buf, err := ioutil.ReadFile("../../../../../../../tools/gif2mp4/safe_image.gif")
	if err != nil {
		logx.Error("open error: %v", err)
		return
	}

	logx.Infof("aaaa", "bb")

	szParts := len(buf) / 10240
	lastSize := len(buf) % 10240
	_ = lastSize
	// idx := sz/10240
	for i := 0; i < szParts; i++ {
		s.dao.WriteFilePartData(context.Background(), 100, 1, int32(i), buf[i*10240:(i+1)*10240])
	}
	logx.Infof("aaaa", "cc")
	if lastSize > 0 {
		logx.Infof("aaaa: lastSize = %d, bufSize = %d", lastSize, len(buf))
		// szParts++
		s.dao.WriteFilePartData(context.Background(), 100, 1, int32(szParts), buf[szParts*10240:])
	}

	logx.Infof("aaaa", "dd")

	if err = s.dao.SetFileInfo(context.Background(), &model.DfsFileInfo{
		Creator:        100,
		FileId:         1,
		Big:            false,
		FileName:       "safe_image.gif",
		FileTotalParts: szParts + 1,
		FilePartSize:   10240,
		// FileSize:       int64(len(buf)),
	}); err != nil {
		logx.Errorf("%v", err)
	}
}

func (s *Server) Destroy() {
	// s.dao.Close()
}

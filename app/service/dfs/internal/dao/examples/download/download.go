package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/commands"

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

	r, err := s.dao.OpenFile(context.Background(), 7997959636588162716, -8695032368284712706, 1)
	if err != nil {
		logx.Error("open error: %v", err)
		// panic(err)
		return
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		logx.Error("open error: %v", err)
		// panic(err)
		return
	}

	b2, err := r.ReadAll(context.Background())
	if err != nil {
		logx.Error("open error: %v", err)
		// panic(err)
		return
	}

	logx.Infof("%s", fmt.Sprintf("%x", md5.Sum(b)))
	logx.Infof("%s", fmt.Sprintf("%x", md5.Sum(b2)))
	err = ioutil.WriteFile("test001.jpeg", b, 0644)
	if err != nil {
		logx.Errorf("write error: %v", err)
	}
}

func (s *Server) Destroy() {
	// s.dao.Close()
}

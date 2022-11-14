// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/teamgram/marmota/pkg/commands"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/config"
	"github.com/teamgram/teamgram-server/app/service/dfs/internal/dao"

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

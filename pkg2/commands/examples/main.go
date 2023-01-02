package main

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/commands"
)

type exapmlesInstance struct {
	state int
	m     func()
}

func (e *exapmlesInstance) Initialize() error {
	logx.SetUp(logx.LogConf{
		ServiceName:         "examples",
		Mode:                "file",
		Encoding:            "json",
		TimeFormat:          "",
		Path:                "logs",
		Level:               "info",
		Compress:            false,
		KeepDays:            0,
		StackCooldownMillis: 100,
		MaxBackups:          0,
		MaxSize:             0,
		Rotation:            "daily",
	})
	logx.Info("null instance initialize")
	return nil
}

func (e *exapmlesInstance) RunLoop() {
	logx.Info("null run_loop...")
	// commands.GSignal <- syscall.SIGQUIT
	commands.DoExit()
}

func (e *exapmlesInstance) Destroy() {
	logx.Info("null destroy...")
}

func main() {
	commands.Run(&exapmlesInstance{})
}

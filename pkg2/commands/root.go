package commands

import (
	"errors"
	"flag"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/version"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// //////////////////////////////////////////////////////////////
var (
	GMainInst MainInstance
	GSignal   chan os.Signal
	ver       = flag.Bool("version", false, "prints current version")
)

type MainInstance interface {
	Initialize() error
	RunLoop()
	Destroy()
}

func Run(inst MainInstance) {
	flag.Parse()

	if *ver {
		v := version.GetVersion()
		fmt.Printf("Version: %s\nGitBranch: %s\nCommitId: %s\nBuild Date: %s\nGo Version: %s\nOS/Arch: %s\n", v.Version, v.GitBranch, v.GitCommit, v.BuildDate, v.GoVersion, v.Platform)
		os.Exit(0)
	}

	defer logx.Close()

	if inst == nil {
		panic(errors.New("inst is nil, exit"))
		return
	}

	//
	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())

	logx.Info("instance initialize...")
	err := inst.Initialize()
	logx.Info("inited")
	if err != nil {
		panic(err)
		return
	}

	// global
	GMainInst = inst

	logx.Info("instance run_loop...")
	go inst.RunLoop()

	GSignal = make(chan os.Signal, 1)
	signal.Notify(GSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-GSignal
		logx.Infof("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			logx.Infof("instance exit...")
			inst.Destroy()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func DoExit() {
	GSignal <- syscall.SIGQUIT
}

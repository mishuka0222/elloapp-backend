package commands

import (
	"sync"
	"testing"

	"github.com/zeromicro/go-zero/core/logx"
)

type NullInstance struct {
	state int
	m     func()
}

func (e *NullInstance) Initialize() error {
	logx.Info("null instance initialize...")
	e.state = 1
	return nil
}

func (e *NullInstance) RunLoop() {
	logx.Info("null run_loop...")
	e.state = 2
	e.m()
}

func (e *NullInstance) Destroy() {
	logx.Info("null destroy...")
	e.state = 3
}

func TestRun1(t *testing.T) {
	instance := &NullInstance{}
	//instance.m = func() {
	//	QuitAppInstance()
	//}
	//
	//DoMainAppInstance(instance)

	result := instance.state
	expect := 3

	if result != expect {
		t.Error(`expect:`, expect, `result:`, result)
	}
}

func TestRun2(t *testing.T) {
	instance := &NullInstance{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	instance.m = func() {
		logx.Info("done...")
		wg.Done()
	}

	// go DoMainAppInstance(instance)
	wg.Wait()
	result := instance.state
	expect := 2

	if result != expect {
		t.Error(`expect:`, expect, `result:`, result)
	}
}

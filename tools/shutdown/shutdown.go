package shutdown

import (
	// "github.com/sirupsen/logrus"
	// "gitlab.baifu-tech.net/dsg-game/backend-server/tools/logs"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"
	"time"
)

var (
	InstanceShutdown *Shutdown
	once             sync.Once
)

// /应用
type Shutdown struct {
	isShutdown bool
	// log           *logrus.Entry
	signalCh      chan os.Signal
	shutdownHooks []ShutdownHook
	sync.RWMutex
}

// /工厂方法
func NewShutdown() *Shutdown {
	once.Do(func() {
		InstanceShutdown = &Shutdown{
			// log:           logs.GetLogger().WithField("module", "Shutdown"),
			signalCh:      make(chan os.Signal, 1),
			shutdownHooks: make([]ShutdownHook, 0),
		}
	})
	return InstanceShutdown
}

func GetShutdown() *Shutdown {
	if InstanceShutdown == nil {
		NewShutdown()
	}
	return InstanceShutdown
}

// /通知关闭
func (object *Shutdown) notifyShutdown() {
	object.Lock()
	sort.Sort(ShutdownHooks(object.shutdownHooks))
	for _, v := range object.shutdownHooks {
		// object.log.Infof("app before shutdown: %s", v.Name())
		v.BeforeShutdown()
	}
	object.Unlock()
}

// /是否已关闭
func (object *Shutdown) IsShutdown() bool {
	return object.isShutdown
}

// /安装关闭钩子
func (object *Shutdown) InstallShutdownHook(hook ShutdownHook) *Shutdown {
	object.Lock()
	if hook != nil {
		object.shutdownHooks = append(object.shutdownHooks, hook)
	}
	object.Unlock()
	return object
}

// /等待关闭
func (object *Shutdown) WaitShutdown() {
	signal.Notify(object.signalCh)
	for s := range object.signalCh {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM:
			// object.log.Infof("Signal: %v, Begin Shutdown", s)
			object.isShutdown = true
			// timeout暫時由docker-compose 控制
			//go object.TimeOut()
			object.notifyShutdown()
			close(object.signalCh)
			// object.log.Info("Shutdown Finished")
			os.Exit(0)
		default:
			//object.log.Infof("signal: %v, not handled", s)
		}
	}
}

// 超時強制關閉
func (object *Shutdown) TimeOut() {
	for {
		select {
		case <-time.After(time.Second * 60):
			// object.log.Error("Shutdown Timeout!!!!!!!!")
			os.Exit(0)
		}
	}
}

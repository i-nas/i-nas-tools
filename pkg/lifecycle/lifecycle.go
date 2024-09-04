package lifecycle

import "github.com/i-nas/i-nas-tools/pkg/log"

var __beforeExitFuncs []func()
var __initFuncs []func()

func RegisterInit(f func()) {
	__initFuncs = append(__initFuncs, f)
}

func RegisterBeforeExit(f func()) {
	__beforeExitFuncs = append(__beforeExitFuncs, f)
}

func BeforeExit() {
	log.Info("start before exit")
	for _, f := range __beforeExitFuncs {
		defer recover()
		f()
	}
	log.Info(" before exit done")
}

func Init() {
	log.Info("start init")
	for _, f := range __initFuncs {
		f()
	}
	log.Info("init done")
}

package system

import (
	"go.uber.org/zap"
	"os"
	"runtime/debug"
	"time"
)

func HandlePanic() {
	// Каждую горутина имеет свой собственный стек, поэтому recover() не может уронить приложения.
	// У горутины свой стек. Паника раскручивает стек до обнаружения обработчика или выйдет из приложения.
	if err := recover(); err != nil {
		zap.L().Error(
			"handle panic",
			zap.Any("recover", err),
			zap.ByteString("stack", debug.Stack()),
		)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
}
func Go(run func()) {
	go func() {
		//Паника в goroutine может уронить приложение.
		defer HandlePanic()
		run()
	}()
}

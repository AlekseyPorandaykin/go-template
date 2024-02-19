package main

import (
	"github.com/AlekseyPorandaykin/go-template/cmd"
	"github.com/AlekseyPorandaykin/go-template/pkg/logger"
	"github.com/AlekseyPorandaykin/go-template/pkg/metrics"
	"github.com/AlekseyPorandaykin/go-template/pkg/shutdown"
	"go.uber.org/zap"
)

var version string

func main() {
	logger.InitDefaultLogger()
	defer func() { _ = zap.L().Sync() }()
	zap.L().Debug("Start app", zap.String("version", version))
	go func() {
		defer shutdown.HandlePanic()
		if err := metrics.Handler("localhost", "9089"); err != nil {
			zap.L().Fatal("error start metric", zap.Error(err))
		}
	}()
	cmd.Execute()
}

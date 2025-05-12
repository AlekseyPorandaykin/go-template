package main

import (
	"github.com/AlekseyPorandaykin/go-template/cmd"
	"github.com/AlekseyPorandaykin/go-template/pkg/logger"
	"github.com/AlekseyPorandaykin/go-template/pkg/monitoring"
	"github.com/AlekseyPorandaykin/go-template/pkg/system"
	"go.uber.org/zap"
)

var version string

func main() {
	logger.InitDefaultLogger()
	defer func() { _ = zap.L().Sync() }()
	zap.L().Debug("Start app", zap.String("version", version))
	system.Go(func() {
		if err := monitoring.Handler("localhost", "9089"); err != nil {
			zap.L().Fatal("error start metric", zap.Error(err))
		}
	})
	cmd.Execute()
}

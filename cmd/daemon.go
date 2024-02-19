package cmd

import (
	"context"
	"time"

	"github.com/AlekseyPorandaykin/go-template/pkg/shutdown"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run daemon script",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go func() {
			defer shutdown.HandlePanic()
			zap.L().Info("Start empty daemon")
			time.Sleep(5 * time.Second)
			zap.L().Info("Stop empty daemon")
			cancel()
		}()

		<-ctx.Done()
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}

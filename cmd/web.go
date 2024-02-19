package cmd

import (
	"context"
	"errors"
	"net"

	"github.com/AlekseyPorandaykin/go-template/pkg/server/http"
	"github.com/AlekseyPorandaykin/go-template/pkg/shutdown"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run web server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		s := http.NewServer()
		defer s.Close()
		go func() {
			defer shutdown.HandlePanic()
			defer cancel()
			if err := s.Run("localhost", "8080"); err != nil && !errors.Is(err, net.ErrClosed) {
				zap.L().Error("error run server", zap.Error(err))
			}
		}()
		<-ctx.Done()
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}

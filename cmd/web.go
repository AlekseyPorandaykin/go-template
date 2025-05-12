package cmd

import (
	"context"
	"github.com/AlekseyPorandaykin/go-template/pkg/server/http"
	"github.com/AlekseyPorandaykin/go-template/pkg/system"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"net"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run web server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		s := http.NewServer()
		defer s.Close()
		system.Go(func() {
			defer cancel()
			if err := s.Run("localhost", "8080"); err != nil && !errors.Is(err, net.ErrClosed) {
				zap.L().Error("error run server", zap.Error(err))
			}
		})
		<-ctx.Done()
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}

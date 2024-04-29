package consoles

import (
	"future-admin/internal/routes"
	"future-admin/pkg/env"
	"future-admin/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func WebServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "start static web server",
		Run: func(cmd *cobra.Command, args []string) {
			if env.IsLocal() {
				viper.Set("port", "8080")
			}

			s := server.NewApiServer()

			routes.RegisterRoute(s)

			s.Run()
		},
	}
}

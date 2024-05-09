package cmd

import (
	"future-admin/cmd/consoles"
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{}

func Execute() {
	rootCmd.AddCommand(consoles.WebServerCmd(), consoles.GenCmd, consoles.MigrateCmd)

	carbon.SetDefault(carbon.Default{
		Layout:   carbon.DateTimeMilliLayout,
		Timezone: carbon.Local,
		Locale:   "zh",
	})

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

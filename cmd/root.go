package cmd

import (
	"future-admin/cmd/consoles"
	_ "future-admin/drivers"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cobra"
	"log"
	"runtime"
)

var rootCmd = &cobra.Command{}

func Execute() {
	vips.Startup(&vips.Config{
		ConcurrencyLevel: runtime.NumCPU(),
		MaxCacheFiles:    0,
		MaxCacheMem:      0,
		MaxCacheSize:     0,
		ReportLeaks:      true,
		CacheTrace:       false,
		CollectStats:     false,
	})
	defer vips.Shutdown()

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

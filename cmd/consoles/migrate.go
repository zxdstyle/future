package consoles

import (
	"future-admin/internal/model"
	"future-admin/pkg"
	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database scheme",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.DBManager().DB().AutoMigrate(model.Album{}, model.Storage{}, model.Image{})
	},
}

package consoles

import (
	"future-admin/pkg"
	"github.com/spf13/cobra"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	GenCmd = &cobra.Command{
		Use: "gen",
		Run: func(cmd *cobra.Command, args []string) {
			g := gen.NewGenerator(gen.Config{
				OutPath:           "internal/dao",
				ModelPkgPath:      "internal/model",
				WithUnitTest:      false,
				FieldNullable:     true,
				FieldCoverable:    true,
				FieldSignable:     true,
				FieldWithIndexTag: false,
				FieldWithTypeTag:  true,
				Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
			})

			g.UseDB(pkg.DBManager().DB())

			g.WithDataTypeMap(dataMap)

			g.WithImportPkgPath("github.com/golang-module/carbon/v2")

			var (
				albums   = g.GenerateModel("albums")
				storages = g.GenerateModel("storages",
					gen.FieldType("addition", "string"),
				)
			)

			g.ApplyBasic(albums, storages)

			g.Execute()
		},
	}

	dataMap = map[string]func(gorm.ColumnType) string{
		"datetime": func(columnType gorm.ColumnType) string {
			return "*carbon.Carbon"
		},
	}
)

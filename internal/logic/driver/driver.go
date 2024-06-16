package driver

import (
	"future-admin/internal/model"
	"future-admin/pkg/utils/co"
	"reflect"
	"strings"
)

var (
	drivers    = co.NewMap[string, model.IDriver]()
	driverInfo = co.NewMap[string, model.DriverInfo]()
)

type Constructor func() model.IDriver

func Register(constructor Constructor) {
	driver := constructor()
	cfg := driver.Config()
	drivers.Set(cfg.Slug, driver)

	t := reflect.TypeOf(driver.GetAddition())
	driverInfo.Set(cfg.Slug, model.DriverInfo{
		Name:      cfg.Name,
		Slug:      cfg.Slug,
		Additions: getAdditionalItems(t),
	})
}

func getAdditionalItems(t reflect.Type) []model.DriverAdditionItem {
	var items []model.DriverAdditionItem
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() == reflect.Struct {
			items = append(items, getAdditionalItems(field.Type)...)
			continue
		}
		tag := field.Tag
		ignore, ok1 := tag.Lookup("ignore")
		name, ok2 := tag.Lookup("json")
		if (ok1 && ignore == "true") || !ok2 {
			continue
		}
		item := model.DriverAdditionItem{
			Name:     name,
			Type:     strings.ToLower(field.Type.Name()),
			Default:  tag.Get("default"),
			Options:  tag.Get("options"),
			Required: tag.Get("required") == "true",
			Help:     tag.Get("help"),
		}
		if tag.Get("type") != "" {
			item.Type = tag.Get("type")
		}
		if item.Name == "root_folder_id" || item.Name == "root_folder_path" {
			//if item.Default == "" {
			//	item.Default = defaultRoot
			//}
			item.Required = item.Default != ""
		}
		// set default type to string
		if item.Type == "" {
			item.Type = "string"
		}
		items = append(items, item)
	}
	return items
}

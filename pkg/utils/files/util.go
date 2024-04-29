package files

import (
	"errors"
	"future-admin/internal/model"
	"github.com/dsoprea/go-exif/v3"
	"sort"
	"strings"
)

func SortFiles(files []model.FileDescription) []model.FileDescription {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Filename < files[j].Filename
	})

	result := make([]model.FileDescription, 0)
	for i, file := range files {
		if file.IsDir {
			result = append(result, files[i])
			continue
		}
	}

	for i, file := range files {
		if !file.IsDir {
			result = append(result, files[i])
			continue
		}
	}

	return result
}

func GetFileExif(file string) ([]model.ExifItem, error) {
	raw, err := exif.SearchFileAndExtractExif(file)
	if errors.Is(err, exif.ErrNoExif) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	entries, _, err := exif.GetFlatExifData(raw, nil)
	if err != nil {
		return nil, err
	}

	exifs := make([]model.ExifItem, 0)
	for _, v := range entries {
		item, ok := model.SupportedExifFormats[v.TagName]
		if !ok {
			exifs = append(exifs, model.ExifItem{
				Key:   v.TagName,
				Value: v.FormattedFirst,
				Label: v.TagName,
			})
			continue
		}

		exifs = append(exifs, model.ExifItem{
			Key:   v.TagName,
			Value: v.FormattedFirst,
			Label: item.Label,
			Enums: item.Enums,
		})
	}

	return exifs, nil
}

func IsImage(file string) bool {
	exts := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".svg"}
	for _, ext := range exts {
		if strings.HasSuffix(file, ext) {
			return true
		}
	}
	return false
}

package images

import (
	"future-admin/pkg/log"
	"github.com/davidbyttow/govips/v2/vips"
	"strings"
)

func ConvertToWebp(file string, quality int, stripMetadata bool) ([]byte, error) {
	img, err := vips.LoadImageFromFile(file, &vips.ImportParams{})
	if err != nil {
		return nil, err
	}

	var (
		buffer []byte
	)
	if quality > 100 {
		buffer, _, err = img.ExportWebp(&vips.WebpExportParams{
			Lossless:      true,
			StripMetadata: stripMetadata,
		})
	} else {
		ep := vips.WebpExportParams{
			Quality:       quality,
			Lossless:      false,
			StripMetadata: stripMetadata,
		}
		for i := range 7 {
			ep.ReductionEffort = i
			buffer, _, err = img.ExportWebp(&ep)
			if err != nil && strings.Contains(err.Error(), "unable to encode") {
				log.Warnf("Can't encode image to WebP with ReductionEffort %d, trying higher value...", i)
			} else if err != nil {
				log.Warnf("Can't encode source image to WebP:%v", err)
			} else {
				break
			}
		}
		buffer, _, err = img.ExportWebp(&ep)
	}
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

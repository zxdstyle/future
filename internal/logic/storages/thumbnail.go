package storages

import (
	"fmt"
	"future-admin/internal/model"
	"github.com/davidbyttow/govips/v2/vips"
	"strings"
)

func (i *Instance) Thumbnail(file string) (*model.FileData, error) {
	content, err := i.driver.Read(file)
	if err != nil {
		return nil, err
	}

	data, err := i.generateThumbnail(content)
	if err != nil {
		return nil, err
	}

	elem := strings.Split(file, "/")
	filename := elem[len(elem)-1]
	eles := strings.Split(filename, ".")

	name := eles[0] + ".avif"

	return &model.FileData{
		Ext:  "avif",
		Name: "thumbnails/" + name,
		Data: data,
		Headers: map[string]string{
			"Content-Disposition": fmt.Sprintf(`attachment; filename="%s"`, name),
			"Content-Type":        "image/" + eles[1],
		},
	}, nil
}

func (i *Instance) generateThumbnail(content []byte) ([]byte, error) {
	img, err := vips.NewImageFromBuffer(content)
	if err != nil {
		return nil, err
	}
	if err := img.Thumbnail(800, 800, vips.InterestingAttention); err != nil {
		return nil, err
	}
	avif := vips.NewWebpExportParams()
	data, _, err := img.ExportWebp(avif)
	if err != nil {
		return nil, err
	}
	return data, nil

}

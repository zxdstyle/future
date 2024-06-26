package main

import (
	"fmt"
	"github.com/davidbyttow/govips/v2/vips"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	vips.Startup(nil)
	defer vips.Shutdown()

	img, err := vips.NewImageFromFile("IMG_0171-2.jpg")
	if err != nil {
		t.Fatal(err)
	}

	//avif := vips.NewAvifExportParams()
	//bytes, meta, err := img.ExportAvif(avif)
	//if err != nil {
	//	t.Fatal(err)
	//}

	webp := vips.NewWebpExportParams()
	webp.Quality = 80
	bytes, meta, err := img.ExportWebp(webp)

	os.WriteFile("test.webp", bytes, 0644)
	fmt.Println(meta)
}

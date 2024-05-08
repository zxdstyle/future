package images

import (
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	webp, err := ConvertToWebp("/Users/zxdstyle/Downloads/金华4.5浦东/极简8x8安吉拉相册  25张  11P/10.jpg", 10, false)
	if err != nil {
		t.Fatal(err)
	}
	os.WriteFile("./test.webp", webp, 0600)
}

package tests

import (
	"fmt"
	"future-admin/pkg/utils/images"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	filepath.Walk("/Users/zxdstyle/Documents/baby", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".jpg") {
			webp, err := images.ConvertToWebp(path, 10, false)
			if err != nil {
				return err
			}

			os.WriteFile("/Users/zxdstyle/baby/"+info.Name()+".webp", webp, 0600)
		}

		if info.IsDir() {
			fmt.Println("文件夹：", path)
		} else {
			fmt.Println("文件：", path)
		}
		return nil
	})
	//images.ConvertToWebp()
}

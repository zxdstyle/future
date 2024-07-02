//go:build darwin

package files

import (
	"github.com/golang-module/carbon/v2"
	"os"
	"syscall"
)

func GetFileCreateTime(info os.FileInfo) carbon.Carbon {
	sys := info.Sys()
	stat := sys.(*syscall.Stat_t)
	return carbon.CreateFromTimestamp(stat.Birthtimespec.Sec)
}

func GetFileUpdateTime(info os.FileInfo) carbon.Carbon {
	sys := info.Sys()
	stat := sys.(*syscall.Stat_t)
	return carbon.CreateFromTimestamp(stat.Mtimespec.Sec)
}

func GetFileAccessTime(info os.FileInfo) carbon.Carbon {
	sys := info.Sys()
	stat := sys.(*syscall.Stat_t)
	return carbon.CreateFromTimestamp(stat.Atimespec.Sec)
}

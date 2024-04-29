//go:build linux

package files

import (
	"github.com/golang-module/carbon/v2"
	"os"
	"syscall"
)

func GetFileCreateTime(info os.FileInfo) carbon.Carbon {
	sys := info.Sys()
	stats := sys.(*syscall.Stat_t)
	return carbon.CreateFromTimestamp(stats.Ctim.Sec)

}

func GetFileUpdateTime(info os.FileInfo) carbon.Carbon {
	sys := info.Sys()
	stats := sys.(*syscall.Stat_t)
	return carbon.CreateFromTimestamp(stats.Mtim.Sec)
}

func GetFileAccessTime(info os.FileInfo) carbon.Carbon {
	sys := info.Sys()
	stats := sys.(*syscall.Stat_t)
	return carbon.CreateFromTimestamp(stats.Atim.Sec)
}

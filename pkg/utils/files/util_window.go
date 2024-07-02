//go:build windows

package files

import (
	"github.com/golang-module/carbon/v2"
	"os"
	"syscall"
)

func GetFileCreateTime(info os.FileInfo) carbon.Carbon {
	sys := info.Sys()
	stat := sys.(*syscall.Win32FileAttributeData)
	return carbon.CreateFromTimestamp(stat.CreationTime.Nanoseconds() / 1e9)
}

func GetFileUpdateTime(info os.FileInfo) carbon.Carbon {
	sys := info.Sys()
	stat := sys.(*syscall.Win32FileAttributeData)
	return carbon.CreateFromTimestamp(stat.LastWriteTime.Nanoseconds() / 1e9)
}

func GetFileAccessTime(info os.FileInfo) carbon.Carbon {
	sys := info.Sys()
	stat := sys.(*syscall.Win32FileAttributeData)
	return carbon.CreateFromTimestamp(linuxstatStats.LastAccessTime.Nanoseconds() / 1e9)
}

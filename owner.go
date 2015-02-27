// +build linux || darwin

package main

import (
	"os"
	"syscall"
)

func init() {
	fileOwner = func(fi os.FileInfo) (uid uint32, ok bool) {
		stat_t, ok := fi.Sys().(*syscall.Stat_t)
		if ok {
			return stat_t.Uid, true
		}
		return 0, false
	}
}

package utils

// the code used here was gotten from https://github.com/Teamwork/reload/blob/master/reload.go
// huge thank you teamwork!!
import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

var (
	binSelf string

	// The watcher won't be closed automatically, and the file descriptor will be
	// leaked if we don't close it in Exec(); see #9.
	closeWatcher func() error
)

// Exec replaces the current process with a new copy of itself.
func Exec() {
	execName := binSelf
	if execName == "" {
		selfName, err := self()
		if err != nil {
			panic(fmt.Sprintf("cannot restart: cannot find self: %v", err))
		}
		execName = selfName
	}

	if closeWatcher != nil {
		closeWatcher()
	}

	err := syscall.Exec(execName, append([]string{execName}, os.Args[1:]...), os.Environ())
	if err != nil {
		panic(fmt.Sprintf("cannot restart: %v", err))
	}
}
func self() (string, error) {
	bin := os.Args[0]
	if !filepath.IsAbs(bin) {
		var err error
		bin, err = os.Executable()
		if err != nil {
			return "", fmt.Errorf(
				"cannot get path to binary %q (launch with absolute path): %w",
				os.Args[0], err)
		}
	}
	return bin, nil
}

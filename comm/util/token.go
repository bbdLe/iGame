package util

import (
	"os/exec"
	"runtime"
	"strings"
)

func NewToken() string {
	if runtime.GOOS == "windows" {
		return "windows_test"
	} else {
		output, err := exec.Command("uuidgen").Output()
		if err != nil {
			return ""
		}
		uuid := strings.TrimSuffix(string(output), "\n")
		return uuid
	}
}
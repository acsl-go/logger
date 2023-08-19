package logger

import (
	"io/ioutil"
	"strings"
)

func getProcessName(d string) string {
	cmdlineFile := "/proc/self/cmdline"
	content, err := ioutil.ReadFile(cmdlineFile)
	if err != nil {
		return d
	}

	// The cmdline file contains the command-line arguments separated by null bytes.
	// We split the content by null bytes and take the first element as the process name.
	parts := strings.Split(string(content), "\x00")
	if len(parts) > 0 {
		return parts[0]
	}

	return d
}

package logger

import (
	"fmt"
	"testing"
)

func TestGetProcessName(t *testing.T) {
	fmt.Println("ProcessName:", getProcessName("default"))
}

package perf

import (
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

// DefaultChecker defines the structure of the default performance checker entity
type DefaultChecker struct {
}

var execCommand = exec.Command

// NewDefaultChecker instantiates a default checker
func NewDefaultChecker() *DefaultChecker {
	return new(DefaultChecker)
}

// DiskPerf tests the storage read/write performance
func (c *DefaultChecker) DiskPerf() (info DiskPerfInfo, err error) {
	var data []byte

	exe := execCommand("dd", "bs=1M", "count=256", "if=/dev/zero", "of=test", "conv=fdatasync")
	data, err = exe.CombinedOutput()

	if err == nil {
		out := string(data)
		logrus.Info(out)

		rate := strings.Split(strings.Split(strings.Replace(out, "\n", "", -1), " s, ")[1], " ")

		info = DiskPerfInfo{
			WriteSpeed: rate[0],
			Unit:       rate[1],
		}
	}
	return
}

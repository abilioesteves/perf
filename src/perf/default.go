package perf

import (
	"fmt"
	"os/exec"
	"path"
	"strings"

	"github.com/abilioesteves/goh/gohtypes"
	"github.com/sirupsen/logrus"
)

// DefaultChecker defines the structure of the default performance checker entity
type DefaultChecker struct {
	WritePath string
}

var execCommand = exec.Command

// NewDefaultChecker instantiates a default checker
func NewDefaultChecker(writePath string) *DefaultChecker {
	return &DefaultChecker{
		WritePath: writePath,
	}
}

// DiskPerf tests the storage read/write performance
func (c *DefaultChecker) DiskPerf() (info DiskPerfInfo, err error) {
	var data []byte

	exe := execCommand("dd", "bs=1M", "count=256", "if=/dev/zero", fmt.Sprintf("of=%v", path.Join(c.WritePath, "test")), "conv=fdatasync")
	data, err = exe.CombinedOutput()
	gohtypes.PanicIfError("Not possible to execute the 'dd' command.", 500, err)

	out := string(data)
	logrus.Info(out)

	rate := strings.Split(strings.Split(strings.Replace(out, "\n", "", -1), " s, ")[1], " ")

	info = DiskPerfInfo{
		WriteSpeed: rate[0],
		Unit:       rate[1],
	}
	return
}

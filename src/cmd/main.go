package main

import (
	"flag"
	"fmt"
	"path"

	"github.com/abilioesteves/goh/gohtypes"
	"github.com/abilioesteves/perf/src/hook"
	"github.com/abilioesteves/perf/src/perf"
)

var writePath string

func init() {
	flag.StringVar(&writePath, "wp", "/data", "the absolute path of a directory to write benchmark data to")
	flag.Parse()
	checkFlags()
}

func main() {
	perfChecker := perf.NewDefaultChecker(writePath)
	webhook := hook.NewDefaultHook(perfChecker)
	go webhook.Init() // inits the webhook
	select {}         // keep-alive magic
}

func checkFlags() {
	if !path.IsAbs(writePath) {
		gohtypes.Panic(fmt.Sprintf("wp '%v' invalid; should be absolute.", writePath), 1)
	}
}

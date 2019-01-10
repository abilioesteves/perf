package main

import (
	"github.com/abilioesteves/perf/src/hook"
	"github.com/abilioesteves/perf/src/perf"
)

func main() {
	perfChecker := perf.NewDefaultChecker()
	webhook := hook.NewDefaultHook(perfChecker)
	go webhook.Init() // inits the webhook
	select {}         // keep-alive magic
}

package hook

import "net/http"

// WebHook defines the contracts of the methods that should be implemented by concrete hooks
type WebHook interface {
	Init()
	DiskPerf(w http.ResponseWriter, r *http.Request)
}

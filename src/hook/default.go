package hook

import (
	"net/http"
	"os"

	"github.com/abilioesteves/goh/gohserver"
	"github.com/abilioesteves/goh/gohtypes"
	"github.com/abilioesteves/perf/src/perf"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// DefaultHook defines the structure that implements the API
type DefaultHook struct {
	router      *mux.Router
	perfChecker perf.Checker
}

// NewDefaultHook instantiates a new default hook structure
func NewDefaultHook(perfChecker perf.Checker) *DefaultHook {
	hook := &DefaultHook{
		router:      mux.NewRouter(),
		perfChecker: perfChecker,
	}

	hook.router.HandleFunc("/diskperf", hook.DiskPerf).Methods("GET")

	return hook
}

// Init initializes the webhook
func (hook *DefaultHook) Init() {
	logrus.Info("Initializing the default webhook...")
	err := http.ListenAndServe("0.0.0.0:17333", hook.router)
	if err != nil {
		logrus.Errorf("Error initializing the default webhook: %v", err)
		os.Exit(1)
	}
	logrus.Info("Default webhook initialized!")
}

// DiskPerf Gets the read and write performance of the disk
func (hook *DefaultHook) DiskPerf(w http.ResponseWriter, r *http.Request) {
	defer gohserver.HandleError(w)

	info, err := hook.perfChecker.DiskPerf()
	gohtypes.PanicIfError("Not possible to check disk performance", 500, err)

	gohserver.WriteJSONResponse(info, 200, w)
}

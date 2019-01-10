package hook

import (
	"fmt"
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

	hook.router.HandleFunc("/diskperf/{format}", hook.DiskPerf).Methods("GET")

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
	format := "json"

	vars := mux.Vars(r)
	if len(vars) > 0 {
		format = vars["format"]
	}

	info, err := hook.perfChecker.DiskPerf()
	gohtypes.PanicIfError("Not possible to check disk performance", 500, err)

	if format == "prom" {
		writePrometheusResponse(info, 200, w)
	} else {
		gohserver.WriteJSONResponse(info, 200, w)
	}
}

func writePrometheusResponse(info perf.DiskPerfInfo, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)

	response := fmt.Sprintln("# HELP disk_write_perf A gauge for the disk performance")
	response = response + fmt.Sprintln("# TYPE disk_write_perf gauge")
	response = response + fmt.Sprintf("disk_write_perf{unit=\"%v\"} %v", info.Unit, info.WriteSpeed)
	_, err := w.Write([]byte(response))
	gohtypes.PanicIfError(fmt.Sprintf("Not possible to write %v response", status), 500, err)

	logrus.Infof("200 Response sent. Payload: %s", info)
}

package perf

// Checker defines the interface of a performance checker entity
type Checker interface {
	DiskPerf() (DiskPerfInfo, error)
}

// DiskPerfInfo defines the information in bytes/s of a disk performance check
type DiskPerfInfo struct {
	Read string

	Write string
}

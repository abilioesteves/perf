package perf

// Checker defines the interface of a performance checker entity
type Checker interface {
	DiskPerf() (DiskPerfInfo, error)
}

// DiskPerfInfo defines the information in of a disk performance check
type DiskPerfInfo struct {
	WriteSpeed string

	Unit string
}

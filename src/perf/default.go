package perf

// DefaultChecker defines the structure of the default performance checker entity
type DefaultChecker struct {
}

// NewDefaultChecker instantiates a default checker
func NewDefaultChecker() *DefaultChecker {
	return new(DefaultChecker)
}

// DiskPerf tests the storage read/write performance
func (c *DefaultChecker) DiskPerf() (DiskPerfInfo, error) {
	return DiskPerfInfo{}, nil
}

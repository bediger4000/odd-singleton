package singleton

import (
	"fmt"
	"sync"
)

type OddSingleton struct {
	serialNumber int
}

var instances [2]*OddSingleton
var nextInstance int
var lck sync.Mutex

// GetInstance rather than the problem specified getInstance
// due to Go's package exported identifier rules.
func GetInstance() *OddSingleton {
	var singleton *OddSingleton
	(&lck).Lock()
	defer (&lck).Unlock()

	if instances[nextInstance] == nil {
		instances[nextInstance] = &OddSingleton{serialNumber: nextInstance + 1}
	}
	singleton = instances[nextInstance]
	nextInstance = (nextInstance + 1) % 2
	return singleton
}

func (p *OddSingleton) String() string {
	return fmt.Sprintf("OddSingleton instance %d", p.serialNumber)
}

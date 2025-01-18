package balancer

import (
	"time"

	"golang.org/x/exp/rand"
)

func BalancerBackend(backend []string) (string, bool) {

	if len(backend) == 0 {
		return "", false
	}

	rand.Seed(uint64(time.Now().UnixNano()))
	index := rand.Intn(len(backend))

	return backend[index], true
}

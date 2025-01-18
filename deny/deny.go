package deny

import (
	"slices"
)

func DenyIP(ips []string, remoteIP string) bool {
	return slices.Contains(ips, remoteIP)
}

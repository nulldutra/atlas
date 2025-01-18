package denyip

import (
	"fmt"
	"slices"
)

func DenyIP(ips []string, remoteIP string) bool {
	fmt.Println(ips)
	fmt.Println(remoteIP)

	return slices.Contains(ips, remoteIP)
}

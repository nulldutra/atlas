package deny

import (
	"slices"
	"strings"
)

func DenyIP(ips []string, remoteIP string) bool {
	return slices.Contains(ips, remoteIP)
}

func DenyHTTPHeader(remoteHeaders []string, header string) bool {
	return slices.Contains(remoteHeaders, header)
}

func DenyHTTPBody(body string, denyBody string) bool {
	return strings.Contains(body, denyBody)
}

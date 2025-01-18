package inspect

import (
	"atlas/deny"
	"net/http"
	"strings"
)

type InspectHTTPRequest struct {
	DenyIPList     []string
	DenyHTTPHeader []string
}

func NewInspectHTTPRequest(denyIPList []string) *InspectHTTPRequest {
	return &InspectHTTPRequest{
		DenyIPList: denyIPList,
	}
}

func (i InspectHTTPRequest) DenyIP(r *http.Request) bool {
	remoteAddr := strings.Split(r.RemoteAddr, ":")
	denyIP := deny.DenyIP(i.DenyIPList, remoteAddr[0])

	return denyIP
}

func (i InspectHTTPRequest) DenyHeader(r *http.Request) bool {
	remoteAddr := strings.Split(r.RemoteAddr, ":")
	denyIP := deny.DenyIP(i.DenyIPList, remoteAddr[0])

	return denyIP
}

package inspect

import (
	"atlas/denyip"
	"net/http"
	"strings"
)

type InspectHTTPRequest struct {
	DenyIPList []string
}

func NewInspectHTTPRequest(denyIPList []string) *InspectHTTPRequest {
	return &InspectHTTPRequest{
		DenyIPList: denyIPList,
	}
}

func (i InspectHTTPRequest) DenyIP(r *http.Request) bool {
	remoteAddr := strings.Split(r.RemoteAddr, ":")
	denyIP := denyip.DenyIP(i.DenyIPList, remoteAddr[0])

	return denyIP
}

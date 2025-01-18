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

func NewInspectHTTPRequest(denyIPList, denyHTTPHeader []string) *InspectHTTPRequest {
	return &InspectHTTPRequest{
		DenyIPList:     denyIPList,
		DenyHTTPHeader: denyHTTPHeader,
	}
}

func (i InspectHTTPRequest) DenyIP(r *http.Request) bool {
	remoteAddr := strings.Split(r.RemoteAddr, ":")
	denyIP := deny.DenyIP(i.DenyIPList, remoteAddr[0])

	return denyIP
}

func (i InspectHTTPRequest) DenyHeader(r *http.Request) bool {
	remoteHeaders := r.Header

	for _, remoteHeader := range remoteHeaders {
		for _, denyHeader := range i.DenyHTTPHeader {
			denyHeader := deny.DenyHTTPHeader(remoteHeader, denyHeader)

			if denyHeader {
				return true
			}
		}
	}

	return false
}

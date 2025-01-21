package inspect

import (
	"atlas/deny"
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type InspectHTTPRequest struct {
	DenyIPList     []string
	DenyHTTPHeader []string
	DenyHTTPBody   []string
}

func NewInspectHTTPRequest(denyIPList, denyHTTPHeader, denyHTTPBody []string) *InspectHTTPRequest {
	return &InspectHTTPRequest{
		DenyIPList:     denyIPList,
		DenyHTTPHeader: denyHTTPHeader,
		DenyHTTPBody:   denyHTTPBody,
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

func (i InspectHTTPRequest) DenyBody(body []byte) bool {
	for _, denyBody := range i.DenyHTTPBody {

		rawBody, _ := url.QueryUnescape(string(body))
		denyBody := deny.DenyHTTPBody(rawBody, denyBody)

		if denyBody {
			return true
		}
	}

	return false
}

func (i InspectHTTPRequest) InspectRequest(w http.ResponseWriter, r *http.Request) bool {
	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(body))

	denyIP := i.DenyIP(r)
	denyHTTPHeader := i.DenyHeader(r)
	denyHTTPBody := i.DenyBody(body)

	if denyIP {
		http.Error(w, "Your IP Address is on the deny list.", http.StatusForbidden)
		return true
	}

	if denyHTTPHeader {
		http.Error(w, "Your requests were blocked because you sent unauthorized headers.", http.StatusForbidden)
		return true
	}

	if denyHTTPBody {
		http.Error(w, "Detected malicious requests.", http.StatusForbidden)
		return true
	}

	return false
}

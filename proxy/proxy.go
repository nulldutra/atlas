package proxy

import (
	"atlas/balancer"
	"atlas/inspect"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Proxy struct {
	Backend        []string
	DenyIPList     []string
	DenyHTTPHeader []string
	DenyHTTPBody   []string
}

func NewProxy(backend, denyIPList, denyHTTPHeader, denyHTTPBody []string) *Proxy {
	return &Proxy{
		Backend:        backend,
		DenyIPList:     denyIPList,
		DenyHTTPHeader: denyHTTPHeader,
		DenyHTTPBody:   denyHTTPBody,
	}
}

func (p *Proxy) Server(w http.ResponseWriter, r *http.Request) {
	backend, _ := balancer.BalancerBackend(p.Backend)
	inspect := inspect.NewInspectHTTPRequest(p.DenyIPList, p.DenyHTTPHeader, p.DenyHTTPBody)

	remote, err := url.Parse(backend)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(body))

	denyIP := inspect.DenyIP(r)
	denyHTTPHeader := inspect.DenyHeader(r)
	denyHTTPBody := inspect.DenyBody(body)

	if denyIP {
		http.Error(w, "Your IP Address is on the deny list.", http.StatusForbidden)
		return
	}

	if denyHTTPHeader {
		http.Error(w, "Your requests were blocked because you sent unauthorized headers.", http.StatusForbidden)
		return
	}

	if denyHTTPBody {
		http.Error(w, "Detected malicous requests.", http.StatusForbidden)
		return
	}

	r.Host = remote.Host
	r.URL.Host = remote.Host
	r.URL.Scheme = remote.Scheme
	r.RequestURI = ""

	client := http.Client{
		Timeout: 60 * time.Second,
	}

	resp, err := client.Do(r)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()

	for k, v := range resp.Header {
		w.Header()[k] = v
	}

	w.WriteHeader(resp.StatusCode)

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
	}
}

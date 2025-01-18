package proxy

import (
	"atlas/balancer"
	"atlas/inspect"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Proxy struct {
	Backend    []string
	DenyIPList []string
}

func NewProxy(backend, denyIPList []string) *Proxy {
	return &Proxy{
		Backend:    backend,
		DenyIPList: denyIPList,
	}
}

func (p *Proxy) Server(w http.ResponseWriter, r *http.Request) {
	backend, _ := balancer.BalancerBackend(p.Backend)

	remote, err := url.Parse(backend)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}

	inspect := inspect.InspectRequest(r, p.DenyIPList)

	fmt.Println(inspect)

	endpoint := remote.String() + r.URL.Path

	nextRequest, err := http.NewRequest(r.Method, endpoint, r.Body)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}

	nextRequest.Header = r.Header

	resp, err := http.DefaultClient.Do(nextRequest)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}

	for k, v := range resp.Header {
		w.Header()[k] = v
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(responseBytes)
}

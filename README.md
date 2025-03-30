## Atlas WaF

Is an simple Web Application Firewall entirely created in Go language.
The reason for the project is for study Golang and how handle HTTP requests.
<img align="right" width="310" height="240" src="https://raw.githubusercontent.com/nulldutra/atlas/main/.logo/logo.png">

The proxy server inspects all HTTP requests and blocked requests based on:

* IP Address

* HTTP Headers

* HTTP Body

## Config

Example of the configuration.

* denyIPList: List of blocked IPs

* denyHTTPHeader: blocks requests that contain a certain word in the header

* denyHTTPBody:blocks requests that contain a certain word in the body

* backend: The proxy server backend. Allow multiple backends and create a simple balance between them.

<hr>


```
denyIPList:
  - 192.168.1.1
  - 192.168.1.2
  - 192.168.1.3
  - 192.168.1.4
  - 192.168.1.107
  - 192.168.88.251

denyHTTPHeader:
  - curl/8.9.1

DenyHTTPBody:
    - bash
    - ping
    - ls
    - echo

backend:
  - http://192.168.88.250:80
  - http://192.168.88.251:80
  - http://192.168.88.252:80
```

## Metrics

Another reason for the project is to study about instrumentation with Prometheus.

The metrics can be accessed from the `/metrics` endpoint.

Example: http://localhost:8000/metrics

<hr>

## Build

Go is a dependencie to build the project. If you need to install, see the documentation: https://go.dev/doc/install


### Building

```
make build
```

### Running

```
make run
```

## References

* https://pkg.go.dev/strings

* https://pkg.go.dev/slices

* https://pkg.go.dev/net/http

* https://prometheus.io/docs/guides/go-application/

* https://antonio-cooler.gitbook.io/coolervoid-tavern/waf-from-the-scratch

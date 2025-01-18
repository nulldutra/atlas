package inspect

import (
	"atlas/denyip"
	"fmt"
	"net/http"
	"strings"
)

func InspectRequest(r *http.Request, denyIPList []string) bool {
	remoteAddr := strings.Split(r.RemoteAddr, ":")

	denyIP := denyip.DenyIP(denyIPList, remoteAddr[0])

	fmt.Println(denyIP)

	return denyIP
}

package gohubspot

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// DumpRequest shows the request at console
func DumpRequest(req *http.Request, showBody bool) {
	requestDump, err := httputil.DumpRequest(req, showBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}

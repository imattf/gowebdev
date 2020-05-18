//not sure what to do one this one...
//have to use autocert because the
// "rsc.io/letsencrypt" has been deleted
// so i tried: https://godoc.org/golang.org/x/crypto/acme/autocert#example-Manager
// or https://pkg.go.dev/golang.org/x/crypto/acme/autocert?tab=doc#example-Manager

package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {

	http.HandleFunc("/", foo)

	m := &autocert.Manager{
		Cache:      autocert.DirCache("certs"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("localhost"),
	}
	s := &http.Server{
		Addr:      ":https",
		TLSConfig: m.TLSConfig(),
	}
	s.ListenAndServeTLS("", "")
}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello TLS")
}

// tls-https
// w/ self-signed cert
// dev tls port is 10443 and prod is 443
package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	//http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	http.ListenAndServeTLS(":10443", "localhost.crt", "localhost.key", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "hey ho, on the tls show")
	// w.Header().Set("Content-Type", "text/plain")
	// w.Write([]byte("This is an example server.\n"))
}

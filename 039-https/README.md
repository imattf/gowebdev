 Go to https:localhost:10443/ or https:127.0.0.1:10443/
 list of TCP ports:
 https:en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers

 Generate unsigned certificate
 go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com
 for example
 go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost


Generate a localhost trusted certificate
```
openssl req -x509 -out localhost.crt -keyout localhost.key \
  -newkey rsa:2048 -nodes -sha256 \
  -subj '/CN=localhost' -extensions EXT -config <( \
   printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:localhost\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")
```
You can then configure your local web server with localhost.crt and localhost.key, and install localhost.crt in your list of locally trusted roots

see: https://letsencrypt.org/docs/certificates-for-localhost/
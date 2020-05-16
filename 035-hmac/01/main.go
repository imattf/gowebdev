// hash message authentication code (HMAC) stuff

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getCode("bob@aol.com")
	fmt.Println(c)

	c = getCode("boob@aol.com")
	fmt.Println(c)

}

func getCode(s string) string {
	hash := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(hash, s)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

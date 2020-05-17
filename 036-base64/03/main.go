//base64 encoding
// just using default the encoding standard

package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {

	s := `Love is but a song to sing
Fear's the way we die
You can make the mountains ring
Or make the angels cry
Though the bird is on the wing
And you may not know why
Come on people now
Smile on your brother
Everybody get together
Try to love one another
Right now`

	//s := "Love is but a song to sing"
	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	// fmt.Println(len(s))
	// fmt.Println(len(s64))
	// fmt.Println("=======")
	// fmt.Println(s)
	// fmt.Println("=======")
	fmt.Println(s64)

	fmt.Println("=======")

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatal("All shes got, Captain!")
	}
	fmt.Println(string(bs))
}

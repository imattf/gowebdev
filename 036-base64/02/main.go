//base64 encoding
// just using default the encoding standard

package main

import (
	"encoding/base64"
	"fmt"
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

	//encodedStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	//s64 := base64.NewEncoding(encodedStd).EncodeToString([]byte(s))
	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println("=======")
	fmt.Println(s)
	fmt.Println("=======")
	fmt.Println(s64)

}

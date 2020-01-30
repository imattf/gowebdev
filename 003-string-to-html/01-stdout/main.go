//using std-out

package main

import (
	"fmt"
)

func main() {

	name := "Kon Tiki"

	tmplate :=
		//	`The world is round according to ` + name + ` and he has spoken`
		`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		<h1>` + name + `</h1>
		</body>
		</html>
		`
	fmt.Println(tmplate)
}

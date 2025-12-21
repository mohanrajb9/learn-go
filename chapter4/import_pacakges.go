/* 
copy the package greetings from ../go/src/greetings to your systems GOPATH
if you dont know the GOPATH, simply run this program
go run import_pacakges.go

this will throw the error for ex
import_pacakges.go:6:2: package greetings is not in std (C:\Program Files\Go\src\greetings)

which means you have to copy the package to C:\Program Files\Go\src\greetings

*/

package main

import (
    "greetings"
	"fmt"
)

func main() {
	const TriangeSides int = 3
	const RectangeSides = 4
	fmt.Println(TriangeSides, RectangeSides)
	greetings.Hi()
	greetings.Hello()
}

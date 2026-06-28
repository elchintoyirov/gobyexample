// [_Buyruq qatori argumentlari_](https://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// dasturlarning bajarilishini parametrlashning keng
// tarqalgan usulidir. Masalan, `go run hello.go` `go`
// dasturiga `run` va `hello.go` argumentlarini ishlatadi.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` xom buyruq qatori argumentlariga kirish
	// imkonini beradi. E'tibor bering, bu slice'dagi
	// birinchi qiymat dasturga yo'l (path) bo'ladi,
	// `os.Args[1:]` esa dasturga argumentlarni saqlaydi.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// Alohida argumentlarni oddiy indekslash orqali olishingiz
	// mumkin.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

// `go` vositasi yoki `git` kabi ba'zi buyruq qatori
// vositalari ko'plab *subbuyruqlar*ga ega bo'ladi, ularning
// har biri o'z bayroqlari to'plamiga ega. Masalan,
// `go build` va `go get` `go` vositasining ikkita turli
// subbuyrug'idir. `flag` paketi bizga o'z bayroqlariga ega
// oddiy subbuyruqlarni osongina aniqlash imkonini beradi.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Biz `NewFlagSet` funksiyasi yordamida subbuyruqni e'lon
	// qilamiz va ushbu subbuyruqqa xos yangi bayroqlarni
	// aniqlashga o'tamiz.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// Boshqa subbuyruq uchun biz turli qo'llab-quvvatlanadigan
	// bayroqlarni aniqlashimiz mumkin.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Subbuyruq dasturning birinchi argumenti sifatida
	// kutiladi.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Qaysi subbuyruq chaqirilganini tekshiramiz.
	switch os.Args[1] {

	// Har bir subbuyruq uchun biz uning o'z bayroqlarini
	// tahlil qilamiz va oxiridagi pozitsion argumentlarga
	// kirish imkoniga ega bo'lamiz.
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

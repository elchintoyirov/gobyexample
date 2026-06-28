// Standart kutubxonaning `strings` paketi satrlar bilan bog'liq
// ko'plab foydali funksiyalarni taqdim etadi. Mana paket haqida
// tasavvur berish uchun ba'zi misollar.

package main

import (
	"fmt"
	s "strings"
)

// Biz quyida uni ko'p ishlatganimiz uchun `fmt.Println` ga
// qisqaroq nom (taxallus) beramiz.
var p = fmt.Println

func main() {

	// Mana `strings` da mavjud bo'lgan funksiyalarning namunasi.
	// Bular satr obyektining metodlari emas, balki paketdagi
	// funksiyalar bo'lgani uchun, biz kerakli satrni funksiyaga
	// birinchi argument sifatida uzatishimiz kerak. Yana ko'plab
	// funksiyalarni [`strings`](https://pkg.go.dev/strings)
	// paketi hujjatlarida topishingiz mumkin.
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}

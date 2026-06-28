// `for` Go'dagi yagona sikl konstruksiyasidir. Mana
// `for` sikllarining ba'zi asosiy turlari.

package main

import "fmt"

func main() {

	// Eng asosiy tur, bitta shart bilan.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Klassik boshlang'ich/shart/keyingi `for` sikli.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// "Buni N marta bajar" degan asosiy iteratsiyani amalga
	// oshirishning yana bir usuli butun son ustidan `range`
	// qilishdir.
	for i := range 3 {
		fmt.Println("range", i)
	}

	// Shartsiz `for` sikldan `break` qilib chiqmaguningizcha
	// yoki o'rab turuvchi funksiyadan `return` qilmaguningizcha
	// takror-takror takrorlanadi.
	for {
		fmt.Println("loop")
		break
	}

	// Shuningdek, siklning keyingi iteratsiyasiga `continue`
	// qilishingiz mumkin.
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// Go'da `if` va `else` bilan tarmoqlanish oddiy.

package main

import "fmt"

func main() {

	// Mana oddiy misol.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Else'siz `if` operatoriga ega bo'lishingiz mumkin.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// `&&` va `||` kabi mantiqiy operatorlar ko'pincha
	// shartlarda foydali bo'ladi.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// Operator shartlardan oldin kelishi mumkin; ushbu
	// operatorda e'lon qilingan har qanday o'zgaruvchilar
	// joriy va keyingi barcha tarmoqlarda mavjud bo'ladi.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// E'tibor bering, Go'da shartlar atrofida qavslar kerak
// emas, lekin jingalak qavslar talab qilinadi.

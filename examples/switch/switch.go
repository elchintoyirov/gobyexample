// _Switch operatorlari_ ko'plab tarmoqlar bo'ylab shartlarni
// ifodalaydi.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Mana oddiy `switch`.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Bitta `case` operatorida bir nechta ifodalarni ajratish
	// uchun vergullardan foydalanishingiz mumkin. Bu misolda biz
	// ixtiyoriy `default` holatidan ham foydalanamiz.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Ifodasiz `switch` if/else mantiqini ifodalashning muqobil
	// usulidir. Bu yerda biz `case` ifodalari konstanta bo'lmasligi
	// mumkinligini ham ko'rsatamiz.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Tip `switch`'i qiymatlar o'rniga tiplarni solishtiradi.
	// Siz buni interfeys qiymatining tipini aniqlash uchun
	// ishlatishingiz mumkin. Bu misolda `t` o'zgaruvchisi o'z
	// bandiga mos keladigan tipga ega bo'ladi.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

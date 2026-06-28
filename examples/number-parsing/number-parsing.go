// Satrlardan sonlarni tahlil qilish ko'plab dasturlarda
// oddiy lekin keng tarqalgan vazifa; mana buni Go da qanday
// bajarish mumkin.

package main

// Ichki o'rnatilgan `strconv` paketi sonni tahlil qilishni
// ta'minlaydi.
import (
	"fmt"
	"strconv"
)

func main() {

	// `ParseFloat` da bu `64` qancha bit aniqlikni tahlil
	// qilishni bildiradi.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// `ParseInt` uchun `0` asosni satrdan aniqlashni
	// bildiradi. `64` esa natija 64 bitga sig'ishini talab qiladi.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` o'n oltilik (hex) formatdagi sonlarni taniydi.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// `ParseUint` ham mavjud.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` oddiy 10 lik asosdagi `int` ni tahlil qilish uchun
	// qulay funksiya.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Tahlil qilish funksiyalari yaroqsiz kiritmada xato qaytaradi.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}

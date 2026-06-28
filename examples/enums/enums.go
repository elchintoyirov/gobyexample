// _Sanab o'tilgan tiplar_ (enumlar) [sum
// tiplar](https://en.wikipedia.org/wiki/Algebraic_data_type)ning maxsus holatidir.
// Enum bu cheklangan miqdordagi mumkin bo'lgan qiymatlarga ega tip bo'lib,
// ularning har biri alohida nomga ega. Go'da enum tipi alohida til
// xususiyati sifatida mavjud emas, biroq enumlarni mavjud til
// idiomalaridan foydalanib amalga oshirish oson.

package main

import "fmt"

// Bizning `ServerState` enum tipimiz asosida `int` tipiga ega.
type ServerState int

// `ServerState` uchun mumkin bo'lgan qiymatlar konstantalar
// sifatida aniqlanadi. Maxsus kalit so'z [iota](https://go.dev/ref/spec#Iota)
// ketma-ket konstanta qiymatlarni avtomatik ravishda hosil qiladi; bu
// holatda 0, 1, 2 va hokazo.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// [fmt.Stringer](https://pkg.go.dev/fmt#Stringer) interfeysini amalga
// oshirish orqali `ServerState` qiymatlarini chop etish yoki satrlarga
// aylantirish mumkin.
//
// Agar mumkin bo'lgan qiymatlar ko'p bo'lsa, bu noqulay bo'lib qolishi
// mumkin. Bunday holatlarda jarayonni avtomatlashtirish uchun [stringer
// vositasi](https://pkg.go.dev/golang.org/x/tools/cmd/stringer) `go:generate`
// bilan birgalikda ishlatilishi mumkin. Batafsil tushuntirish uchun [bu
// post](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)ni ko'ring.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// Agar bizda `int` tipidagi qiymat bo'lsa, uni `transition` ga uzata
	// olmaymiz - kompilyator tip mosligi yo'qligidan shikoyat qiladi. Bu
	// enumlar uchun ma'lum darajada kompilyatsiya vaqtidagi tip xavfsizligini ta'minlaydi.

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition server uchun holat o'tishini emulyatsiya qiladi;
// u mavjud holatni qabul qiladi va yangi holatni qaytaradi.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Faraz qilaylik, keyingi holatni aniqlash uchun bu yerda
		// biror predikatlarni tekshiramiz...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

// Go'da xatolarni oshkora, alohida qaytariladigan qiymat orqali
// uzatish idiomatik hisoblanadi. Bu Java, Python va Ruby kabi
// tillarda ishlatiladigan istisnolardan (exceptions) va C'da ba'zan
// ishlatiladigan ortiqcha yuklatilgan yagona natija / xato qiymatidan
// farq qiladi. Go'ning yondashuvi qaysi funksiyalar xato qaytarishini
// ko'rishni va ularni boshqa, xato bo'lmagan vazifalar uchun
// qo'llaniladigan til konstruksiyalari bilan qayta ishlashni osonlashtiradi.
//
// Qo'shimcha tafsilotlar uchun [errors paketi](https://pkg.go.dev/errors)
// hujjatlarini va [ushbu blog postni](https://go.dev/blog/go1.13-errors) ko'ring.

package main

import (
	"errors"
	"fmt"
)

// Kelishuvga ko'ra, xatolar oxirgi qaytariladigan qiymat bo'lib,
// o'rnatilgan interfeys bo'lgan `error` tipiga ega.
func f(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` berilgan xato xabari bilan oddiy `error`
		// qiymatini quradi.
		return -1, errors.New("can't work with 42")
	}

	// Xato pozitsiyasidagi `nil` qiymat hech qanday xato
	// bo'lmaganligini bildiradi.
	return arg + 3, nil
}

// Sentinel xato oldindan e'lon qilingan o'zgaruvchi bo'lib,
// muayyan xato holatini bildirish uchun ishlatiladi.
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		// Kontekst qo'shish uchun xatolarni yuqori darajadagi xatolar
		// bilan o'rashimiz mumkin. Buni qilishning eng oddiy usuli
		// `fmt.Errorf` dagi `%w` verbidir. O'ralgan xatolar mantiqiy
		// zanjir hosil qiladi (A B'ni o'raydi, u esa C'ni o'raydi va hokazo)
		// va bu zanjirni `errors.Is` hamda `errors.AsType` kabi funksiyalar
		// bilan so'rab olish mumkin.
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// `if` qatorida ichki (inline) xato tekshiruvidan foydalanish
		// idiomatik hisoblanadi.
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

			// `errors.Is` berilgan xato (yoki uning zanjiridagi har qanday xato)
			// muayyan xato qiymatiga mos kelishini tekshiradi. Bu, ayniqsa, o'ralgan
			// yoki ichma-ich joylashgan xatolar bilan foydalidir va sizga xatolar
			// zanjiridagi muayyan xato tiplari yoki sentinel xatolarni aniqlash imkonini beradi.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}

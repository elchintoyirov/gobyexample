// Maxsus xato tiplarini ularda `Error()` metodini
// amalga oshirish orqali aniqlash mumkin. Quyida argument
// xatosini oshkora ifodalash uchun maxsus tipdan foydalanadigan
// yuqoridagi misolning bir variantini ko'rasiz.

package main

import (
	"errors"
	"fmt"
)

// Maxsus xato tipi odatda "Error" qo'shimchasiga ega bo'ladi.
type argError struct {
	arg     int
	message string
}

// Ushbu `Error` metodini qo'shish `argError` ni `error`
// interfeysini amalga oshiradigan qiladi.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {

		// Maxsus xatomizni qaytaramiz.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// `errors.AsType` `errors.Is` ning yanada ilg'or versiyasidir.
	// U berilgan xato (yoki uning zanjiridagi har qanday xato)
	// muayyan xato tipiga mos kelishini tekshiradi va uni o'sha tipdagi
	// qiymatga aylantiradi hamda `true` ni qaytaradi. Agar moslik bo'lmasa,
	// ikkinchi qaytariladigan qiymat `false` bo'ladi.
	_, err := f(42)
	if ae, ok := errors.AsType[*argError](err); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}

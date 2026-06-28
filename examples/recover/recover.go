// Go `recover` o'rnatilgan funksiyasidan foydalanib,
// panikadan _tiklash_ imkonini beradi. `recover` `panic`ning
// dasturni to'xtatishidan saqlab, uning o'rniga bajarilishni
// davom ettirishga imkon beradi.

// Bu qayerda foydali bo'lishiga misol: server mijoz
// ulanishlaridan biri kritik xato chiqarsa, ishdan chiqishni
// xohlamaydi. Buning o'rniga, server o'sha ulanishni yopib,
// boshqa mijozlarga xizmat ko'rsatishni davom ettirishni
// xohlaydi. Aslida, Go'ning `net/http`si HTTP serverlar uchun
// standart holatda shunday qiladi.

package main

import "fmt"

// Bu funksiya panika qiladi.
func mayPanic() {
	panic("a problem")
}

func main() {
	// `recover` kechiktirilgan funksiya ichida chaqirilishi
	// kerak. O'rab turuvchi funksiya panika qilganda, defer
	// faollashadi va uning ichidagi `recover` chaqiruvi
	// panikani ushlaydi.
	defer func() {
		if r := recover(); r != nil {
			// `recover`ning qaytaradigan qiymati `panic`
			// chaqiruvida ko'tarilgan xatodir.
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// Bu kod ishlamaydi, chunki `mayPanic` panika qiladi.
	// `main`ning bajarilishi panika nuqtasida to'xtaydi va
	// kechiktirilgan closure ichida qayta tiklanadi.
	fmt.Println("After mayPanic()")
}

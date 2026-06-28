// Berilgan status bilan darhol chiqish uchun `os.Exit` dan
// foydalaning.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Exit` dan foydalanganda `defer`lar _ishga tushmaydi_,
	// shuning uchun bu `fmt.Println` hech qachon chaqirilmaydi.
	defer fmt.Println("!")

	// 3-status bilan chiqamiz.
	os.Exit(3)
}

// E'tibor bering, masalan C'dan farqli o'laroq, Go chiqish
// statusini bildirish uchun `main` dan butun sonli qaytariladigan
// qiymatdan foydalanmaydi. Agar noldan farqli status bilan
// chiqmoqchi bo'lsangiz, `os.Exit` dan foydalanishingiz kerak.

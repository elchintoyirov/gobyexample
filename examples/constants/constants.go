// Go belgi, satr, boolean va son qiymatlari uchun
// _konstanta_larni qo'llab-quvvatlaydi.

package main

import (
	"fmt"
	"math"
)

// `const` konstanta qiymatni e'lon qiladi.
const s string = "constant"

func main() {
	fmt.Println(s)

	// `const` ifodasi funksiya tanasi ichida ham
	// paydo bo'lishi mumkin.
	const n = 500000000

	// Konstanta ifodalari ixtiyoriy aniqlikda
	// arifmetik amallarni bajaradi.
	const d = 3e20 / n
	fmt.Println(d)

	// Sonli konstanta unga tip berilmaguncha (masalan,
	// oshkora konvertatsiya orqali) hech qanday tipga ega bo'lmaydi.
	fmt.Println(int64(d))

	// Songa tipni uni tip talab qiladigan kontekstda
	// ishlatish orqali berish mumkin, masalan o'zgaruvchiga
	// qiymat berish yoki funksiya chaqiruvi orqali. Masalan, bu yerda
	// `math.Sin` `float64` tipini kutadi.
	fmt.Println(math.Sin(n))
}

// `panic` odatda nimadir kutilmaganda noto'g'ri ketganini
// bildiradi. Asosan biz undan oddiy ish jarayonida yuz
// bermasligi kerak bo'lgan yoki biz to'g'ri ishlashga tayyor
// bo'lmagan xatolarda tezda to'xtash uchun foydalanamiz.

package main

import (
	"os"
	"path/filepath"
)

func main() {

	// Biz bu saytda kutilmagan xatolarni tekshirish uchun
	// panic dan foydalanamiz. Bu saytda panic qilish uchun
	// mo'ljallangan yagona dastur.
	panic("a problem")

	// panic ning keng tarqalgan ishlatilishi - funksiya biz
	// qanday ishlashni bilmaydigan (yoki istamaydigan) xato
	// qiymatini qaytarsa, ishni to'xtatishdir. Mana yangi fayl
	// yaratishda kutilmagan xato olsak, `panic` qilishga misol.
	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}
}

// `filepath` paketi *fayl yo'llarini* operatsion tizimlar
// o'rtasida portativ bo'lgan tarzda tahlil qilish va
// yaratish uchun funksiyalarni taqdim etadi; masalan,
// Linuxda `dir/file`, Windowsda esa `dir\file`.
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// Yo'llarni portativ tarzda yaratish uchun `Join`
	// ishlatilishi kerak. U istalgan miqdordagi argumentni
	// qabul qiladi va ulardan ierarxik yo'l yaratadi.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// `/` yoki `\` belgilarini qo'lda biriktirish o'rniga
	// har doim `Join` dan foydalanishingiz kerak.
	// Portativlikni ta'minlashdan tashqari, `Join` ortiqcha
	// ajratgichlar va katalog o'zgarishlarini olib tashlab,
	// yo'llarni normallashtiradi.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` va `Base` yo'lni katalog va faylga ajratish
	// uchun ishlatilishi mumkin. Muqobil ravishda, `Split`
	// ikkalasini ham bitta chaqiruvda qaytaradi.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// Yo'l absolyut ekanligini tekshirishimiz mumkin.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Ba'zi fayl nomlarida nuqtadan keyin kengaytma bo'ladi.
	// Bunday nomlardan kengaytmani `Ext` bilan ajratib
	// olishimiz mumkin.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Faylning kengaytmasi olib tashlangan nomini topish
	// uchun `strings.TrimSuffix` dan foydalaning.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` *baza* va *maqsad* o'rtasidagi nisbiy yo'lni
	// topadi. Agar maqsadni bazaga nisbatan qilib bo'lmasa,
	// xato qaytaradi.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}

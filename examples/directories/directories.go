// Go fayl tizimidagi *kataloglar* bilan ishlash uchun
// bir nechta foydali funksiyaga ega.

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Joriy ishchi katalogda yangi quyi-katalog yaratamiz.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// Vaqtinchalik kataloglar yaratayotganda ularni o'chirishni
	// `defer` qilish yaxshi amaliyot hisoblanadi. `os.RemoveAll`
	// butun katalog daraxtini o'chiradi (`rm -rf` ga o'xshash).
	defer os.RemoveAll("subdir")

	// Yangi bo'sh fayl yaratish uchun yordamchi funksiya.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// Biz `MkdirAll` bilan ota-kataloglarni ham o'z ichiga olgan
	// kataloglar iyerarxiyasini yaratishimiz mumkin. Bu buyruq
	// satridagi `mkdir -p` ga o'xshash.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` katalog tarkibini ro'yxatlaydi va `os.DirEntry`
	// obyektlari slice'ini qaytaradi.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` joriy ishchi katalogni o'zgartirishga imkon beradi,
	// `cd` ga o'xshash.
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Endi *joriy* katalogni ro'yxatlaganda `subdir/parent/child`
	// tarkibini ko'ramiz.
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Boshlagan joyimizga `cd` qilib qaytamiz.
	err = os.Chdir("../../..")
	check(err)

	// Biz katalogni *rekursiv* tarzda, uning barcha
	// quyi-kataloglari bilan birga aylanib chiqishimiz ham mumkin.
	// `WalkDir` har bir tashrif buyurilgan fayl yoki katalogni
	// qayta ishlash uchun callback funksiyani qabul qiladi.
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
	check(err)
}

// `visit` `filepath.WalkDir` tomonidan rekursiv ravishda
// topilgan har bir fayl yoki katalog uchun chaqiriladi.
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}

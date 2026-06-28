// Go da fayllarga yozish biz oldin o'qish uchun
// ko'rganlarimizga o'xshash usullarga amal qiladi.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Boshlash uchun, satrni (yoki shunchaki baytlarni)
	// faylga qanday yozish quyida keltirilgan.
	d1 := []byte("hello\ngo\n")
	path1 := filepath.Join(os.TempDir(), "dat1")
	err := os.WriteFile(path1, d1, 0644)
	check(err)

	// Yanada batafsil yozishlar uchun faylni yozish uchun oching.
	path2 := filepath.Join(os.TempDir(), "dat2")
	f, err := os.Create(path2)
	check(err)

	// Faylni ochgandan so'ng darhol `Close` ni defer
	// qilish odatiy (idiomatik) hisoblanadi.
	defer f.Close()

	// Kutilganidek, bayt slice larini `Write` qila olasiz.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// `WriteString` ham mavjud.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Yozishlarni barqaror xotiraga o'tkazish uchun `Sync` chaqiring.
	f.Sync()

	// `bufio` biz oldin ko'rgan buferlangan o'quvchilarga
	// qo'shimcha ravishda buferlangan yozuvchilarni ham taqdim etadi.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Barcha buferlangan operatsiyalar asosiy yozuvchiga
	// qo'llanilganligiga ishonch hosil qilish uchun `Flush` dan foydalaning.
	w.Flush()

}

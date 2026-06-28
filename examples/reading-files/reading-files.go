// Fayllarni o'qish va yozish ko'plab Go dasturlari uchun
// kerak bo'ladigan asosiy vazifalardir. Avval fayllarni
// o'qishning ba'zi misollarini ko'rib chiqamiz.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Fayllarni o'qish ko'pchilik chaqiruvlarni xatolarga
// tekshirishni talab qiladi. Bu yordamchi funksiya quyidagi
// xatolarni tekshirishimizni soddalashtiradi.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Ehtimol, eng asosiy fayl o'qish vazifasi faylning butun
	// mazmunini xotiraga yuklashdir.
	path := filepath.Join(os.TempDir(), "dat")
	dat, err := os.ReadFile(path)
	check(err)
	fmt.Print(string(dat))

	// Ko'pincha faylning qaysi qismlari va qanday o'qilishi
	// ustidan ko'proq nazorat xohlaysiz. Bu vazifalar uchun
	// `os.File` qiymatini olish maqsadida faylni `Open` qilishdan
	// boshlang.
	f, err := os.Open(path)
	check(err)

	// Faylning boshidan bir nechta bayt o'qiymiz. 5 tagacha
	// o'qishga ruxsat beramiz, lekin amalda nechtasi o'qilganini
	// ham qayd qilamiz.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Shuningdek, fayldagi ma'lum bir joyga `Seek` qilib, o'sha
	// yerdan `Read` qilishingiz mumkin.
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// Boshqa seek qilish usullari joriy kursor pozitsiyasiga
	// nisbatan,
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	// va faylning oxiriga nisbatan amalga oshiriladi.
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	// `io` paketi fayl o'qish uchun foydali bo'lishi mumkin
	// bo'lgan ba'zi funksiyalarni taqdim etadi. Masalan,
	// yuqoridagiga o'xshash o'qishlar `ReadAtLeast` bilan
	// yanada ishonchli amalga oshirilishi mumkin.
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// O'rnatilgan rewind yo'q, lekin `Seek(0, io.SeekStart)`
	// buni amalga oshiradi.
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// `bufio` paketi buferlangan o'quvchini amalga oshiradi,
	// bu ko'plab kichik o'qishlardagi samaradorligi uchun ham,
	// taqdim etadigan qo'shimcha o'qish metodlari tufayli ham
	// foydali bo'lishi mumkin.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Ishingiz tugagach, faylni yoping (odatda bu `Open`
	// qilingandan so'ng darhol `defer` bilan rejalashtiriladi).
	f.Close()
}

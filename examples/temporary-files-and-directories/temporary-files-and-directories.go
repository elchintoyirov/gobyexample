// Dastur bajarilishi davomida biz ko'pincha dastur chiqib
// ketgandan keyin kerak bo'lmaydigan ma'lumotlarni yaratishni
// xohlaymiz. *Vaqtinchalik fayllar va kataloglar* shu maqsadda
// foydalidir, chunki ular vaqt o'tishi bilan fayl tizimini
// ifloslantirmaydi.

package main

import (
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

	// Vaqtinchalik fayl yaratishning eng oson yo'li
	// `os.CreateTemp` ni chaqirishdir. U fayl yaratadi *va* uni
	// o'qish va yozish uchun ochadi. Biz birinchi argument
	// sifatida `""` beramiz, shuning uchun `os.CreateTemp` faylni
	// OS'imiz uchun sukut bo'yicha joyda yaratadi.
	f, err := os.CreateTemp("", "sample")
	check(err)

	// Vaqtinchalik faylning nomini ko'rsating. Unix asosidagi
	// OS'larda katalog ehtimol `/tmp` bo'ladi. Fayl nomi
	// `os.CreateTemp` ga ikkinchi argument sifatida berilgan
	// prefiks bilan boshlanadi, qolgan qismi esa bir vaqtning
	// o'zidagi chaqiruvlar har doim turli xil fayl nomlarini
	// yaratishini ta'minlash uchun avtomatik tanlanadi.
	fmt.Println("Temp file name:", f.Name())

	// Ishimiz tugagandan keyin faylni tozalang. OS ehtimol bir
	// muncha vaqtdan keyin vaqtinchalik fayllarni o'zi tozalaydi,
	// lekin buni oshkora bajarish yaxshi amaliyotdir.
	defer os.Remove(f.Name())

	// Biz faylga bir oz ma'lumot yozishimiz mumkin.
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// Agar biz ko'plab vaqtinchalik fayllar yozmoqchi bo'lsak,
	// vaqtinchalik *katalog* yaratishni afzal ko'rishimiz mumkin.
	// `os.MkdirTemp` ning argumentlari `CreateTemp` ningidek, lekin
	// u ochiq fayl o'rniga katalog *nomini* qaytaradi.
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// Endi biz vaqtinchalik fayl nomlarini ularning oldiga
	// vaqtinchalik katalogimizni qo'shib hosil qilishimiz mumkin.
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}

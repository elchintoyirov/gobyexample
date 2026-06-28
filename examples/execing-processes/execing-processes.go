// Oldingi misolda biz [tashqi jarayonlarni ishga
// tushirishni](spawning-processes) ko'rib chiqdik. Buni biz
// ishlayotgan Go jarayoniga tashqi jarayon kerak bo'lganda qilamiz.
// Ba'zan biz joriy Go jarayonini butunlay boshqa (balki Go bo'lmagan)
// jarayon bilan almashtirmoqchi bo'lamiz. Buni qilish uchun biz Go'ning
// klassik
// <a href="https://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>
// funksiyasi amalga oshirilishidan foydalanamiz.

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// Misolimiz uchun biz `ls` ni exec qilamiz. Go biz bajarmoqchi
	// bo'lgan binar faylga absolyut yo'lni talab qiladi, shuning uchun
	// biz uni topish uchun `exec.LookPath` dan foydalanamiz (ehtimol
	// `/bin/ls`).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` argumentlarni slice ko'rinishida talab qiladi (bitta
	// katta satr emas). Biz `ls` ga bir nechta keng tarqalgan
	// argumentlarni beramiz. E'tibor bering, birinchi argument dastur
	// nomi bo'lishi kerak.
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` shuningdek foydalanish uchun [muhit
	// o'zgaruvchilari](environment-variables) to'plamiga muhtoj. Bu yerda
	// biz shunchaki joriy muhitimizni taqdim etamiz.
	env := os.Environ()

	// Bu yerda haqiqiy `syscall.Exec` chaqiruvi. Agar bu chaqiruv
	// muvaffaqiyatli bo'lsa, jarayonimizning bajarilishi shu yerda
	// tugaydi va `/bin/ls -a -l -h` jarayoni bilan almashtiriladi.
	// Agar xato bo'lsa, biz qaytariladigan qiymatni olamiz.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

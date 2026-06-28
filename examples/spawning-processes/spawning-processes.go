// Ba'zan Go dasturlarimiz boshqa jarayonlarni ishga
// tushirishi kerak bo'ladi.

package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// Hech qanday argument yoki kirish olmaydigan va shunchaki
	// stdoutga biror narsa chop etadigan oddiy buyruq bilan
	// boshlaymiz. `exec.Command` yordamchisi bu tashqi
	// jarayonni ifodalash uchun obyekt yaratadi.
	dateCmd := exec.Command("date")

	// `Output` metodi buyruqni ishga tushiradi, uning tugashini
	// kutadi va uning standart chiqishini to'playdi. Agar
	// xatolar bo'lmasa, `dateOut` sana ma'lumotini o'z ichiga
	// olgan baytlarni saqlaydi.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// `Output` va `Command`ning boshqa metodlari buyruqni
	// bajarishda muammo bo'lsa (masalan, noto'g'ri yo'l)
	// `*exec.Error`ni, va buyruq ishlagan, lekin nolga teng
	// bo'lmagan qaytarish kodi bilan chiqqan bo'lsa
	// `*exec.ExitError`ni qaytaradi.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		if e, ok := errors.AsType[*exec.Error](err); ok {
			fmt.Println("failed executing:", e)
		} else if e, ok := errors.AsType[*exec.ExitError](err); ok {
			exitCode := e.ExitCode()
			fmt.Println("command exit rc =", exitCode)
		} else {
			panic(err)
		}
	}

	// Keyin biz tashqi jarayonning `stdin`iga ma'lumotlarni
	// uzatadigan va uning `stdout`idan natijalarni
	// yig'adigan biroz murakkabroq holatni ko'rib chiqamiz.
	grepCmd := exec.Command("grep", "hello")

	// Bu yerda kirish/chiqish pipe'larini oshkora olamiz,
	// jarayonni ishga tushiramiz, unga biroz kirish yozamiz,
	// natijaviy chiqishni o'qiymiz va nihoyat jarayonning
	// chiqishini kutamiz.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// Yuqoridagi misolda xatolarni tekshirishni tashlab
	// ketdik, lekin ularning barchasi uchun odatdagi
	// `if err != nil` patternidan foydalanishingiz mumkin.
	// Shuningdek, biz faqat `StdoutPipe` natijalarini yig'amiz,
	// lekin `StderrPipe`ni ham xuddi shunday yig'ishingiz
	// mumkin.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// E'tibor bering, buyruqlarni ishga tushirganda bitta
	// buyruq qatori satrini berishimiz emas, balki oshkora
	// ajratilgan buyruq va argument massivini berishimiz
	// kerak. Agar satr bilan to'liq buyruqni ishga tushirishni
	// xohlasangiz, `bash`ning `-c` opsiyasidan foydalanishingiz
	// mumkin:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

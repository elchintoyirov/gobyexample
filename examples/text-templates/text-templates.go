// Go `text/template` paketi yordamida dinamik kontent yaratish
// yoki foydalanuvchiga moslashtirilgan chiqishni ko'rsatish uchun
// o'rnatilgan qo'llab-quvvatlashni taklif qiladi. `html/template`
// deb nomlangan qarindosh paket xuddi shu API'ni taqdim etadi,
// lekin qo'shimcha xavfsizlik imkoniyatlariga ega va HTML
// generatsiya qilish uchun ishlatilishi kerak.

package main

import (
	"os"
	"text/template"
)

func main() {

	// Biz yangi shablon yaratishimiz va uning tanasini satrdan
	// tahlil qilishimiz mumkin.
	// Shablonlar statik matn va kontentni dinamik tarzda
	// kiritish uchun ishlatiladigan `{{...}}` ichiga olingan
	// "amal"larning aralashmasidir.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// Muqobil ravishda, biz `Parse` xato qaytargan holda panika
	// qilish uchun `template.Must` funksiyasidan foydalanishimiz
	// mumkin. Bu ayniqsa global qamrovda ishga tushirilgan
	// shablonlar uchun foydalidir.
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// Shablonni "bajarish" orqali biz uning matnini amallari
	// uchun aniq qiymatlar bilan hosil qilamiz. `{{.}}` amali
	// `Execute` ga parametr sifatida uzatilgan qiymat bilan
	// almashtiriladi.
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// Biz quyida ishlatadigan yordamchi funksiya.
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// Agar ma'lumot struct bo'lsa, biz uning maydonlariga
	// murojaat qilish uchun `{{.FieldName}}` amalidan
	// foydalanishimiz mumkin. Shablon bajarilayotganda murojaat
	// qilish mumkin bo'lishi uchun maydonlar eksport qilingan
	// bo'lishi kerak.
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// Xuddi shu narsa map'larga ham tegishli; map'larda kalit
	// nomlarining harf registriga cheklov yo'q.
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// if/else shablonlar uchun shartli bajarishni ta'minlaydi.
	// Qiymat tipning sukut qiymati bo'lsa, masalan 0, bo'sh satr,
	// nil ko'rsatkich va hokazo, u false deb hisoblanadi.
	// Bu namuna shablonlarning yana bir imkoniyatini namoyish
	// etadi: bo'sh joyni kesish uchun amallarda `-` dan
	// foydalanish.
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// range bloklari bizga slice'lar, massivlar, map'lar yoki
	// kanallar bo'ylab aylanish imkonini beradi. range bloki
	// ichida `{{.}}` iteratsiyaning joriy elementiga o'rnatiladi.
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}

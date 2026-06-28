// Go tiplarning yanada uzluksiz _kompozitsiyasini_ ifodalash
// uchun struct va interfeyslarni _ichiga joylash_ni
// qo'llab-quvvatlaydi. Buni fayl va papkalarni dastur
// binariga joylash uchun Go 1.16+ versiyasida kiritilgan go
// direktivasi bo'lgan [`//go:embed`](embed-directive) bilan
// adashtirmaslik kerak.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// `container` o'z ichiga `base` ni _joylaydi_. Ichiga joylash
// nomsiz maydonga o'xshaydi.
type container struct {
	base
	str string
}

func main() {

	// Literallar bilan struct yaratishda biz ichiga joylashni
	// oshkora ishga tushirishimiz kerak; bu yerda ichiga
	// joylangan tip maydon nomi vazifasini bajaradi.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// Biz base'ning maydonlariga to'g'ridan-to'g'ri `co` orqali
	// murojaat qila olamiz, masalan `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Muqobil ravishda, biz ichiga joylangan tip nomidan
	// foydalanib to'liq yo'lni yozishimiz mumkin.
	fmt.Println("also num:", co.base.num)

	// `container` o'z ichiga `base` ni joylagani uchun, `base`
	// ning metodlari ham `container` ning metodlariga aylanadi.
	// Bu yerda biz `base` dan ichiga joylangan metodni
	// to'g'ridan-to'g'ri `co` orqali chaqiramiz.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Metodli struct'larni ichiga joylash boshqa struct'larga
	// interfeys implementatsiyalarini berish uchun ishlatilishi
	// mumkin. Bu yerda biz `container` endi `describer`
	// interfeysini amalga oshirishini ko'ramiz, chunki u o'z
	// ichiga `base` ni joylaydi.
	var d describer = co
	fmt.Println("describer:", d.describe())
}

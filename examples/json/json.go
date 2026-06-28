// Go ichki va maxsus ma'lumotlar tiplaridan va ularga
// kodlash hamda dekodlashni qo'shgan holda, JSON kodlash va
// dekodlash uchun ichki qo'llab-quvvatlashni taqdim etadi.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Biz quyida maxsus tiplarni kodlash va dekodlashni
// namoyish qilish uchun ushbu ikkita structdan
// foydalanamiz.
type response1 struct {
	Page   int
	Fruits []string
}

// JSON'da faqat eksport qilingan maydonlar
// kodlanadi/dekodlanadi. Eksport qilinishi uchun maydonlar
// bosh harf bilan boshlanishi kerak.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// Avval biz asosiy ma'lumotlar tiplarini JSON
	// satrlariga kodlashni ko'rib chiqamiz. Mana atomik
	// qiymatlar uchun ba'zi misollar.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// Va mana slicelar va maplar uchun ba'zilari, ular
	// kutganingizdek JSON massivlari va obyektlariga
	// kodlanadi.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON paketi sizning maxsus ma'lumotlar tiplaringizni
	// avtomatik ravishda kodlay oladi. U kodlangan natijaga
	// faqat eksport qilingan maydonlarni kiritadi va standart
	// holatda o'sha nomlarni JSON kalitlari sifatida
	// ishlatadi.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Kodlangan JSON kalit nomlarini sozlash uchun struct
	// maydon e'lonlarida teglardan foydalanishingiz mumkin.
	// Bunday teglarning misolini ko'rish uchun yuqoridagi
	// `response2` ta'rifini tekshiring.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Endi JSON ma'lumotlarini Go qiymatlariga dekodlashni
	// ko'rib chiqaylik. Mana umumiy ma'lumotlar tuzilmasi
	// uchun misol.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Biz JSON paketi dekodlangan ma'lumotlarni joylashtirishi
	// mumkin bo'lgan o'zgaruvchini taqdim etishimiz kerak. Bu
	// `map[string]interface{}` satrlardan ixtiyoriy ma'lumotlar
	// tiplariga bo'lgan mapni saqlaydi.
	var dat map[string]interface{}

	// Mana haqiqiy dekodlash va bog'liq xatolarni tekshirish.
	// Qisqalik uchun biz ushbu misollarda xatolarni
	// e'tiborsiz qoldiramiz; haqiqiy kodda siz doimo xatolarni
	// tekshirishingiz va ularga qarab harakat qilishingiz
	// kerak.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Dekodlangan mapdagi qiymatlardan foydalanish uchun biz
	// ularni mos tipiga aylantirishimiz kerak. Masalan, bu
	// yerda biz `num` dagi qiymatni kutilgan `float64` tipiga
	// aylantiramiz.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Ichma-ich joylashgan ma'lumotlarga kirish bir qator
	// aylantirishlarni talab qiladi.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Shuningdek, biz JSON ni maxsus ma'lumotlar tiplariga
	// dekodlashimiz mumkin. Bu dasturlarimizga qo'shimcha
	// tip xavfsizligini qo'shish va dekodlangan ma'lumotlarga
	// kirishda tip tasdiqlariga bo'lgan ehtiyojni bartaraf
	// etish afzalliklariga ega.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	_ = json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// Yuqoridagi misollarda biz har doim ma'lumotlar va
	// standart chiqishdagi JSON ko'rinishi o'rtasida bayt va
	// satrlardan oraliq sifatida foydalandik. Shuningdek, biz
	// JSON kodlashlarini to'g'ridan-to'g'ri `os.Stdout` kabi
	// `os.Writer` larga yoki hatto HTTP javob tanalariga ham
	// stream qilishimiz mumkin.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	_ = enc.Encode(d)

	// `os.Stdin` kabi `os.Reader` lardan yoki HTTP so'rov
	// tanalaridan stream o'qish `json.Decoder` bilan amalga
	// oshiriladi.
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	_ = dec.Decode(&res1)
	fmt.Println(res1)
}

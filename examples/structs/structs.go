// Go'ning _struct_'lari maydonlarning tiplangan to'plamidir.
// Ular yozuvlarni hosil qilish uchun ma'lumotlarni birga
// guruhlashda foydalidir.

package main

import "fmt"

// Bu `person` struct tipi `name` va `age` maydonlariga ega.
type person struct {
	name string
	age  int
}

// `newPerson` berilgan nom bilan yangi person struct'ini tuzadi.
func newPerson(name string) *person {
	// Go axlatni yig'uvchili (garbage collected) til; siz lokal
	// o'zgaruvchiga ko'rsatkichni xavfsiz qaytarishingiz mumkin -
	// u faqat unga faol murojaatlar qolmaganda axlat yig'uvchi
	// tomonidan tozalanadi.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// Bu sintaksis yangi struct yaratadi.
	fmt.Println(person{"Bob", 20})

	// Struct'ni ishga tushirishda maydonlarni nomlashingiz mumkin.
	fmt.Println(person{name: "Alice", age: 30})

	// Tashlab ketilgan maydonlar nol qiymatga ega bo'ladi.
	fmt.Println(person{name: "Fred"})

	// `&` prefiksi struct'ga ko'rsatkichni beradi.
	fmt.Println(&person{name: "Ann", age: 40})

	// Yangi struct yaratishni konstruktor funksiyalarda qamrab olish odatiy holdir
	fmt.Println(newPerson("Jon"))

	// Struct maydonlariga nuqta orqali murojaat qiling.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Siz nuqtalardan struct ko'rsatkichlari bilan ham
	// foydalanishingiz mumkin - ko'rsatkichlar avtomatik ravishda
	// dereferens qilinadi.
	sp := &s
	fmt.Println(sp.age)

	// Struct'lar o'zgaruvchan (mutable).
	sp.age = 51
	fmt.Println(sp.age)

	// Agar struct tipi faqat bitta qiymat uchun ishlatilsa, biz
	// unga nom berishimiz shart emas. Qiymat anonim struct tipiga
	// ega bo'lishi mumkin. Bu uslub ko'pincha
	// [jadvalga asoslangan testlar](testing-and-benchmarking)
	// uchun ishlatiladi.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}

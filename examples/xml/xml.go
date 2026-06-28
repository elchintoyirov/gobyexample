// Go `encoding/xml` paketi orqali XML va XML ga o'xshash
// formatlar uchun o'rnatilgan qo'llab-quvvatlashni taklif qiladi.

package main

import (
	"encoding/xml"
	"fmt"
)

// Plant XML ga moslashtiriladi. JSON misollariga o'xshab,
// maydon teglari enkoder va dekoder uchun ko'rsatmalarni
// o'z ichiga oladi. Bu yerda XML paketining ba'zi maxsus
// xususiyatlaridan foydalanamiz: `XMLName` maydon nomi bu
// struct ni ifodalovchi XML elementining nomini belgilaydi;
// `id,attr` esa `Id` maydoni ichki element emas, balki XML
// _atributi_ ekanligini bildiradi.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// O'simligimizni ifodalovchi XML ni chiqaramiz;
	// inson uchun o'qish osonroq bo'lgan natija hosil
	// qilish uchun `MarshalIndent` dan foydalanamiz.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// Natijaga umumiy XML sarlavhasini qo'shish uchun
	// uni aniq tarzda qo'shing.
	fmt.Println(xml.Header + string(out))

	// XML li baytlar oqimini ma'lumotlar tuzilmasiga
	// tahlil qilish uchun `Unmarshal` dan foydalaning.
	// Agar XML noto'g'ri shakllangan bo'lsa yoki Plant ga
	// moslashtirib bo'lmasa, tavsiflovchi xato qaytariladi.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// `parent>child>plant` maydon tegi enkoderga barcha
	// `plant` larni `<parent><child>...` ostida joylashtirishni aytadi
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}

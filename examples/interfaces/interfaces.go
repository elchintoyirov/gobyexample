// _Interfeyslar_ - bu metod imzolarining nomlangan
// to'plamlaridir.

package main

import (
	"fmt"
	"math"
)

// Mana geometrik shakllar uchun asosiy interfeys.
type geometry interface {
	area() float64
	perim() float64
}

// Misolimiz uchun biz ushbu interfeysni `rect` va `circle`
// tiplarida amalga oshiramiz.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Go'da interfeysni amalga oshirish uchun biz shunchaki
// interfeysdagi barcha metodlarni amalga oshirishimiz
// kerak. Bu yerda biz `rect` larda `geometry` ni amalga
// oshiramiz.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// `circle` lar uchun amalga oshirish.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Agar o'zgaruvchi interfeys tipiga ega bo'lsa, biz
// nomlangan interfeysdagi metodlarni chaqira olamiz. Mana
// istalgan `geometry` ustida ishlash uchun bundan
// foydalanadigan generik `measure` funksiyasi.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Ba'zan interfeys qiymatining runtime tipini bilish
// foydali bo'ladi. Bir variant - bu yerda ko'rsatilganidek
// *tip tasdig'i* dan foydalanish; boshqasi esa [tip
// `switch`](switch).
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// `circle` va `rect` struct tiplari ikkalasi ham
	// `geometry` interfeysini amalga oshiradi, shuning
	// uchun biz ushbu structlarning namunalarini `measure`
	// ga argument sifatida ishlatishimiz mumkin.
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}

// Go'ning `math/rand/v2` paketi
// [psevdotasodifiy son](https://en.wikipedia.org/wiki/Pseudorandom_number_generator)
// generatsiyasini ta'minlaydi.

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// Masalan, `rand.IntN` tasodifiy `int` n ni qaytaradi,
	// `0 <= n < 100`.
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// `rand.Float64` `float64` `f` ni qaytaradi,
	// `0.0 <= f < 1.0`.
	fmt.Println(rand.Float64())

	// Bundan boshqa oraliqlarda tasodifiy float lar
	// generatsiya qilish uchun foydalanish mumkin, masalan
	// `5.0 <= f' < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Agar sizga ma'lum seed kerak bo'lsa, yangi
	// `rand.Source` yarating va uni `New` konstruktoriga
	// uzating. `NewPCG` ikkita `uint64` sondan iborat seed
	// talab qiladigan yangi
	// [PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator)
	// manbasini yaratadi.
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}

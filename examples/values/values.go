// Go da satrlar, butun sonlar, kasr sonlar, mantiqiy
// qiymatlar va hokazo kabi turli xil qiymat tiplari
// mavjud. Quyida bir nechta oddiy misol keltirilgan.

package main

import "fmt"

func main() {

	// Satrlar, ularni `+` bilan birlashtirish mumkin.
	fmt.Println("go" + "lang")

	// Butun sonlar va kasr sonlar.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Mantiqiy qiymatlar, kutilganidek mantiqiy operatorlar bilan.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

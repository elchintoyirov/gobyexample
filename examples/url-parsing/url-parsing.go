// URL lar [resurslarni topishning yagona usulini](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/) taqdim etadi.
// Quyida Go da URL larni qanday tahlil qilish ko'rsatilgan.

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// Sxema, autentifikatsiya ma'lumotlari, host, port, path,
	// so'rov parametrlari va so'rov fragmentini o'z ichiga
	// olgan ushbu misol URL ni tahlil qilamiz.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// URL ni tahlil qilamiz va xatolar yo'qligiga ishonch hosil qilamiz.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Sxemaga murojaat qilish oddiy.
	fmt.Println(u.Scheme)

	// `User` barcha autentifikatsiya ma'lumotlarini o'z
	// ichiga oladi; alohida qiymatlar uchun bunda
	// `Username` va `Password` ni chaqiring.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host` mavjud bo'lsa, ham hostname ni ham portni o'z
	// ichiga oladi. Ularni ajratib olish uchun
	// `SplitHostPort` dan foydalaning.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Bu yerda `path` ni va `#` dan keyingi fragmentni
	// ajratib olamiz.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// So'rov parametrlarini `k=v` formatidagi satr ko'rinishida
	// olish uchun `RawQuery` dan foydalaning. So'rov
	// parametrlarini map ga ham tahlil qilishingiz mumkin.
	// Tahlil qilingan so'rov parametri map lari satrlardan
	// satrlar slice lariga aylantiriladi, shuning uchun agar
	// faqat birinchi qiymatni istasangiz, `[0]` indeksiga
	// murojaat qiling.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}

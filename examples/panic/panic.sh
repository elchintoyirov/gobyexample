# Bu dasturni ishga tushirish uni panic qildiradi, xato
# xabari va goroutina izlarini chop etadi va nolga teng
# bo'lmagan status bilan chiqadi.

# `main` dagi birinchi panic ishga tushganda, dastur kodning
# qolgan qismiga yetmasdan chiqadi. Agar dastur vaqtinchalik
# fayl yaratishga urinishini ko'rmoqchi bo'lsangiz, birinchi
# panic ni izohga oling.
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# E'tibor bering, ko'p xatolarni qayta ishlash uchun
# istisnolardan (exceptions) foydalanadigan ba'zi tillardan
# farqli o'laroq, Go da imkon qadar xatoni bildiruvchi
# qaytariladigan qiymatlardan foydalanish idiomatikdir.

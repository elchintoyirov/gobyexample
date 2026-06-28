# E'tibor bering, slice'lar massivlardan boshqa tip
# bo'lsa-da, ular `fmt.Println` tomonidan o'xshash tarzda
# tasvirlanadi.
$ go run slices.go
uninit: [] true true
emp: [  ] len: 3 cap: 3
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]

# Go'da slice'larning dizayni va amalga oshirilishi haqida
# batafsil ma'lumot uchun Go jamoasining bu [ajoyib blog
# postini](https://go.dev/blog/slices-intro) ko'rib chiqing.

# Endi massivlar va slice'larni ko'rganimizdan so'ng, Go'ning
# yana bir asosiy o'rnatilgan ma'lumotlar tuzilmasi: maplarni
# ko'rib chiqamiz.

# Dasturni ishga tushirish dasturda o'rnatgan `FOO` qiymatini
# qabul qilishimizni, ammo `BAR` bo'sh ekanligini ko'rsatadi.
$ go run environment-variables.go
FOO: 1
BAR: 

# Muhitdagi kalitlar ro'yxati sizning aniq mashinangizga
# bog'liq bo'ladi.
TERM_PROGRAM
PATH
SHELL
...
FOO

# Agar avval muhitda `BAR` ni o'rnatsak, ishlayotgan dastur
# o'sha qiymatni qabul qiladi.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...

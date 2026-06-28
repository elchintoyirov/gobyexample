# Buyruq qatori bayroqlari dasturi bilan tajriba o'tkazish
# uchun, avval uni kompilyatsiya qilib, so'ngra hosil
# bo'lgan binar faylni to'g'ridan-to'g'ri ishga tushirish
# ma'qul.
$ go build command-line-flags.go

# Qurilgan dasturni avval barcha bayroqlar uchun qiymatlar
# berib sinab ko'ring.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# E'tibor bering, agar bayroqlarni tashlab ketsangiz, ular
# avtomatik ravishda o'zlarining standart qiymatlarini
# oladi.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# Oxiridagi pozitsion argumentlarni istalgan bayroqlardan
# keyin berish mumkin.
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# E'tibor bering, `flag` paketi barcha bayroqlar pozitsion
# argumentlardan oldin kelishini talab qiladi (aks holda
# bayroqlar pozitsion argumentlar sifatida talqin
# qilinadi).
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# Buyruq qatori dasturi uchun avtomatik hosil qilingan
# yordam matnini olish uchun `-h` yoki `--help`
# bayroqlaridan foydalaning.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Agar `flag` paketiga ko'rsatilmagan bayroqni bersangiz,
# dastur xato xabarini chop etadi va yordam matnini yana
# ko'rsatadi.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...

# Dasturni ishga tushirish hashni hisoblaydi va uni
# odam o'qiy oladigan hex formatida chop etadi.
$ go run sha256-hashes.go
sha256 this string
1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a...


# Yuqorida ko'rsatilganga o'xshash usuldan foydalanib boshqa
# hashlarni hisoblashingiz mumkin. Masalan, SHA512 hashlarini
# hisoblash uchun `crypto/sha512`ni import qiling va
# `sha512.New()`dan foydalaning.

# E'tibor bering, agar sizga kriptografik xavfsiz hashlar
# kerak bo'lsa, [hash kuchini](https://en.wikipedia.org/wiki/Cryptographic_hash_function)
# diqqat bilan o'rganishingiz kerak!

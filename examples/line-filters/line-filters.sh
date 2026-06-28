# Qator filtrimizni sinab ko'rish uchun, avval bir nechta
# kichik harfli qatorlardan iborat fayl yarating.
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# So'ngra bosh harfli qatorlarni olish uchun qator
# filtridan foydalaning.
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER

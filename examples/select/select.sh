# Biz kutilganidek `"one"` va keyin `"two"` qiymatlarini
# qabul qilamiz.
$ time go run select.go 
received one
received two

# E'tibor bering, umumiy bajarilish vaqti faqat ~2 sekund,
# chunki 1 va 2 sekundli `Sleep`larning ikkalasi ham
# parallel bajariladi.
real	0m2.245s

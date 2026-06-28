# Ishga tushirilgan dasturlar ularni to'g'ridan-to'g'ri buyruq
# qatoridan ishga tushirgandagidek bir xil chiqish qaytaradi.
$ go run spawning-processes.go 
> date
Thu 05 May 2022 10:10:12 PM PDT

# date'da `-x` flagi yo'q, shuning uchun u xato xabari va
# nolga teng bo'lmagan qaytarish kodi bilan chiqadi.
command exit rc = 1
> grep hello
hello grep

> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go

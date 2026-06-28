#  Agar `exit.go` ni `go run` yordamida ishga tushirsangiz,
# chiqish `go` tomonidan qabul qilinadi va chop etiladi.
$ go run exit.go
exit status 3

# Binar faylni build qilib bajarish orqali statusni
# terminalda ko'rishingiz mumkin.
$ go build exit.go
$ ./exit
$ echo $?
3

# E'tibor bering, dasturimizdagi `!` hech qachon chop etilmadi.

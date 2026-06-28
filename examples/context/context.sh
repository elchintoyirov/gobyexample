# Serverni fonda ishga tushiramiz.
$ go run context.go &

# `/hello` ga mijoz so'rovini simulyatsiya qilamiz va
# bekor qilishni signal qilish uchun boshlanganidan
# ko'p o'tmay Ctrl+C ni bosamiz.
$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended

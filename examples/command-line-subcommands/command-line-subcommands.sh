$ go build command-line-subcommands.go 

# Avval foo subbuyrug'ini chaqiramiz.
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# Endi bar'ni sinab ko'ramiz.
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# Ammo bar foo'ning bayroqlarini qabul qilmaydi.
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# Keyingi misolda dasturlarni parametrlashning yana bir
# keng tarqalgan usuli bo'lgan muhit o'zgaruvchilarini
# ko'rib chiqamiz.

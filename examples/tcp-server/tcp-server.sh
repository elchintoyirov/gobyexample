# TCP serverni fonda ishga tushiring.
$ go run tcp-server.go &

# netcat yordamida ma'lumot yuboring va javobni qabul qiling.
$ echo "Hello from netcat" | nc localhost 8090
ACK: HELLO FROM NETCAT


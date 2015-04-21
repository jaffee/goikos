package main

import ()

func main() {

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func listen() {
	l, err := net.Listen("tcp", "12121")
	check(err)
	for {
		conn, err := l.Accept()
		check(err)
	}
}

func handleConn(conn net.Conn) {
	bytes = make([]byte, 1000)
	// TODO figure out how to read in variable amounts of bytes properly
}

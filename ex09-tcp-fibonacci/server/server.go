package main

import (
	"fmt"
	"math/big"
	"net"
	"strconv"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":4545")
	m := make(map[string]*big.Int)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()
	fmt.Println("Server is listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn, m)
	}
}

func handleConnection(conn net.Conn, m map[string]*big.Int) {
	defer conn.Close()
	for {
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)

		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}

		source := string(input[0:n])
		val, ok := m[source]
		if source == "quit" {
			return
		}
		if ok {
			start1 := time.Now()
			fmt.Println(source, "-", val)
			finish1 := time.Since(start1)
			conn.Write([]byte(finish1.String() + "\n" + val.String()))
		} else {
			start := time.Now()
			target := Fib(source).String()
			finish := time.Since(start)
			m[source] = Fib(source)
			fmt.Println(source, "-", target)
			conn.Write([]byte(finish.String() + "\n" + target))
		}
	}
}

func Fib(n string) *big.Int {
	num, _ := strconv.Atoi(n)
	if num == 0 {
		return big.NewInt(0)
	}
	if num == 1 {
		return big.NewInt(1)
	}
	f1 := Fib(strconv.Itoa(num - 1))
	f2 := Fib(strconv.Itoa(num - 2))
	fib := big.NewInt(0)
	fib.Add(f1, f2)
	return fib
}

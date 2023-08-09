package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

/*
Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123


Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout

*/

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Fprintln(os.Stderr, "usage: go-telnet [--timeout timeout] host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(host, port))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error resolving address: %v\n", err)
		os.Exit(1)
	}

	conn, err := net.DialTimeout("tcp", addr.String(), *timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error reading from server: %v\n", err)
				os.Exit(1)
			}
			fmt.Print(string(buf[:n]))
		}
	}()

	for {
		buf := make([]byte, 1024)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading from stdin: %v\n", err)
			os.Exit(1)
		}
		if n == 0 {
			break
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error writing to server: %v\n", err)
			os.Exit(1)
		}
	}
}

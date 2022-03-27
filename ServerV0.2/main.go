package main

import "fmt"

func main() {
	fmt.Println("程序已经执行")
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}

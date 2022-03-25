package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

//创建一个Server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (this *Server) Handele(conn net.Conn) {
	fmt.Println("开始业务")
}

func (this *Server) Start() {
	// 1.listener
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("连接出错")
	}
	// 2.defer close
	defer listen.Close()

	for {
		// 3.accpect
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("conn failed")
		}
		// 4.doHandel
		go this.Handele(conn)
	}

}

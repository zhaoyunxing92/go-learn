package main

import (
	"bufio"
	"fmt"
	"net"
)

// 网络模块
func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("linsten failed ,err:%v \n", err)
		return
	}

	fmt.Println("tcp server start")

	for {
		//跟客户端建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("conn failed ,err:%v \n", err)
			continue
		}
		go process(conn)
	}

}

func process(conn net.Conn) {
	// 关闭连接
	defer conn.Close()
	for {

		reader := bufio.NewReader(conn)
		var buf [128]byte

		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("server read failed err:%v \n", err)
		}

		recv := string(buf[:n])
		fmt.Println("接受到消息：", recv)
		//给客户端写消息
		_, _ = conn.Write([]byte("ok"))
	}
}

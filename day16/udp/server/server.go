package main

import (
	"fmt"
	"net"
)

// udp
func main() {

	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})

	if err != nil {
		fmt.Printf("server udp listen failad err:%v \n", err)
		return
	}

	defer listen.Close()

	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		fmt.Println("n:", n)
		if err != nil {
			fmt.Printf("server udp read failed err:%v \n", err)
			continue
		}
		fmt.Printf("接收到[%v]的消息:%v \n", addr, string(buf[:n]))

		//给客户端写消息
		_, err = listen.WriteToUDP([]byte("ok"), addr)
		if err != nil {
			fmt.Printf("server write to %v failed,err:%v \n", addr, err)
		}
	}

}

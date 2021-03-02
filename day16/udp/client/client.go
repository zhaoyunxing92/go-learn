package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})

	if err != nil {
		fmt.Printf("dial failed,err:%v \n", err)
		return
	}

	defer conn.Close()

	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入：")
		str, _ := input.ReadString('\n')
		_, err = conn.Write([]byte(str))
		if err != nil {
			fmt.Printf("client to server failed,err:%v \n", err)
			continue
		}

		var buf [1024]byte

		n, addr, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("client read failed,err:%v \n", err)
			continue
		}
		fmt.Printf("服务端[%v]数据:%v \n", addr, string(buf[:n]))
	}

}

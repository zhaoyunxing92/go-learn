package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	dial, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Printf("client dial failed,err:%v \n", err)
		return
	}
	input := bufio.NewReader(os.Stdin)
	fmt.Println("tcp client start")
	for {
		str, _ := input.ReadString('\n')

		str = strings.TrimSpace(str)
		if strings.ToUpper(str) == "Q" {
			return
		}
		//给服务端发消息
		_, err := dial.Write([]byte(str))
		if err != nil {
			fmt.Printf("send failed,err:%v \n", err)
			return
		}
		// 接受服务端数据
		var buf [1024]byte
		n, err := dial.Read(buf[:])
		if err != nil {
			fmt.Printf("client read failad,err:%v \n", err)
		}
		fmt.Println("收到服务端消息：", string(buf[:n]))
	}
}

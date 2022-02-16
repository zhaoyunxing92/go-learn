package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func main() {
	data := [][]string{
		{"org.apache.dubbo.game.basketballService", "CreateUser", "1.0", "dubbo"},
		{"org.apache.dubbo.game.basketballService", "GetUserByCode", "1.0", "dubbo"},
		{"org.apache.dubbo.game.basketballService", "GetUserByName", "1.0", "dubbo"},
		{"org.apache.dubbo.game.basketballService", "GetUserByNameAndAge", "1.0", "dubbo"},
		{"org.apache.dubbo.game.basketballService", "GetUserTimeout", "1.0", "dubbo"},
		{"org.apache.dubbo.game.basketballService", "UpdateUser", "1.0", "dubbo"},
		{"com.dubbogo.pixiu.UserService", "", "1.0", "dubbo"},
		{"com.apache.dubbo.sample.basic.IGreeter", "", "1.0", "dubbo"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"interface", "methods", "version", "group"})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}

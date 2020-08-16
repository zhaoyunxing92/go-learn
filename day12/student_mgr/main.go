package main

import (
	"fmt"
	"os"
)

//显示菜单
func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.显示全部学员信息")
	fmt.Println("4.退出")
}

// 获取用户输入的学员信息
func getInputStudent() *student {
	var (
		id          int
		name, class string
	)
	fmt.Println("请按照要求输入学员信息")

	fmt.Print("请输入学号：")
	_, _ = fmt.Scanf("%d\n", &id)

	fmt.Print("请输入姓名：")
	_, _ = fmt.Scanf("%s\n", &name)

	fmt.Print("请输入班级：")
	_, _ = fmt.Scanf("%s\n", &class)

	return newStudent(id, name, class)
}

// 学员管理系统
func main() {
	showMenu()
	mgr := newStudentMgr()
	for {
		var input int
		//scan获取用户输入
		_, err := fmt.Scanf("%d", &input)
		if err != nil {
			fmt.Println("输入异常，err:", err)
		}
		fmt.Println("用户输入的是：", input)

		switch input {
		case 1:
			mgr.addStudent(getInputStudent())
			break;
		case 2:
			mgr.updateStudent(getInputStudent())
		case 3:
			mgr.showStudents()
		case 4:
			//退出
			os.Exit(0)
		}
	}
}

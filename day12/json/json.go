package main

import (
	"encoding/json"
	"fmt"
)

/**
学生
*/
type Student struct {
	//编号
	Id int `json:"id"`
	//姓名
	Name string `json:"name"`
	//昵称
	NickName string `json:"nickName"`
}

func newStudent(id int, name, nickName string) Student {
	return Student{id, name, nickName}
}

/**
班级
*/
type Class struct {
	// 班级名称
	Title string `json:"title"`
	// 班级学生
	Students []Student `json:"students"`
}

func main() {

	c1 := Class{Title: "304", Students: make([]Student, 0, 5)}

	for i := 0; i < 3; i++ {
		sut := newStudent(i, fmt.Sprintf("stu%02d", i), fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, sut)
	}

	fmt.Printf("%#v\n", c1)
	fmt.Println("=============json marshal=====================")
	//使用json
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed,err:", err)
		return
	}

	fmt.Printf("%s\n", data) //{"title":"304","students":[{"id":0,"name":"stu00"},{"id":1,"name":"stu01"},{"id":2,"name":"stu02"}]}

	fmt.Println("=============json unmarshal=====================")

	//可以发现json字符串首字母大小都可以转换成功,且字段全部大写或小写都可以成功转换
	str := `{"Title":"304","Students":[{"Id":15,"Name":"sunny","nick_name":"云兴"}]}`

	var c2 Class

	err = json.Unmarshal([]byte(str), &c2)

	if err != nil {
		fmt.Println("json unmarshal failed,err:", err)
		return
	}
	fmt.Printf("%#v\n", c2) //main.Class{Title:"304", Students:[]main.Student{main.Student{Id:15, Name:"sunny"}}}
}

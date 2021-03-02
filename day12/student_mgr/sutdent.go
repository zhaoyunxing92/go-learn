package main

import "fmt"

type student struct {
	id          int
	name, class string
}

func newStudent(id int, name, class string) *student {
	return &student{id, name, class}
}

//学员管理
type studentMgr struct {
	allStudents []*student
}

func newStudentMgr() *studentMgr {
	return &studentMgr{allStudents: make([]*student, 0, 10)}
}

//添加学员
func (mgr *studentMgr) addStudent(stu *student) {
	mgr.allStudents = append(mgr.allStudents, stu)
}

//修改学员信息
func (mgr *studentMgr) updateStudent(newStu *student) {
	for idx, st := range mgr.allStudents {
		if st.id == newStu.id {
			mgr.allStudents[idx] = newStu
			return
		}
	}
	fmt.Printf("没有找到学号是【%d】的学生\n", newStu.id)
}

//显示全部信息
func (mgr *studentMgr) showStudents() {
	for _, stu := range mgr.allStudents {
		fmt.Printf("学号:%d 姓名:%s 班级:%s\n", stu.id, stu.name, stu.class)
	}
}

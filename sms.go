package main

import (
	"fmt"
	"os"
)

//学生管理系统 1.数据-->结构体字段 2.功能--->方法

type student struct{
	id int64
	name string
}
//学生的管理者
type studentMgr struct{
	allStudent map[int64]student
}

func showMenu(){
	fmt.Println("welcome to sms")
	fmt.Println(`
	1.查看学生
	2.添加学生
	3.修改学生
	4.删除学生
	5.退出
	`)
}
//查看学生
func (s studentMgr) showStudent(){
	for _,stu:=range s.allStudent{
		fmt.Printf("学号:%d 姓名:%s\n",stu.id,stu.name)
	}
}
//添加学生
func (s studentMgr) addStudent(){
 var (
	 newId int64
	 newName string
 )
 fmt.Print("请输入新加学生的学号:")
 fmt.Scanln(&newId)
 fmt.Print("请输入学生的姓名:")
 fmt.Scanln(&newName)
 newStu:=student{
	 id:newId,
	 name:newName,
 }
 s.allStudent[newStu.id]=newStu
 fmt.Println("添加成功")
}

//修改学生
func (s studentMgr) editStudent(){

}

//删除学生
func (s studentMgr) deleteStudent(){

}

var smg studentMgr

func main(){
	smg=studentMgr{
		allStudent: make(map[int64]student),
	}
	for{
 showMenu()
 var choice int64
 fmt.Print("请输入您的选项:")
 fmt.Scanln(&choice)
 fmt.Printf("您选择的序列是:%d\n",choice)
 switch choice {
 case 1:
 	smg.showStudent()
 case 2:
	smg.addStudent()
 case 3:
	smg.editStudent()
 case 4:
	smg.deleteStudent()
 case 5:
	os.Exit(1)
 default:
	fmt.Println("your Enter is error")

 }
	}
}
package main

import (
	"fmt"
	"log"
)

type ISchoolor interface {
	SetName(name string)
}

type Student struct {
	Name    string
	Age     int32
	ClassNo int32 //班级编号
}

func (p *Student) SetName(name string) {
	p.Name = name
	log.Printf("stu:%p setname:%v\n", p, name)
}

type Teacher struct {
	Name   string
	Age    int32
	Course []string //教课程列表
}

func (p *Teacher) SetName(name string) {
	p.Name = name
	log.Printf("ter:%p setname:%v\n", p, name)
}

func main() {
	fmt.Println("~~~ c12 interface reflect ~~~")

	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	var stu Student = Student{Name: "zhang3", Age: 15, ClassNo: 401}
	var ter Teacher = Teacher{Name: "liu.7", Age: 40, Course: []string{"math", "english"}}
	var ior ISchoolor

	log.Printf("stu:%p %v", &stu, stu)
	log.Printf("ter:%p %v", &ter, ter)
	log.Printf("ior:%p %v", &ior, ior)

	// log.Printf("stu:%v ter:%v ior:%v", stu, ter, ior)
}

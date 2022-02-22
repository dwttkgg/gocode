package nimingstruct

import "fmt"

type Student struct {
	Name  string
	Age   int
	score float64
}

func (stu *Student) testing() {
	fmt.Println("该学生正在考试...")
}

func (stu *Student) SetScore(fl float64) {
	stu.score = fl
}

func (stu *Student) ShowInfo() {
	fmt.Println("该学生的分数是", stu.score)
}

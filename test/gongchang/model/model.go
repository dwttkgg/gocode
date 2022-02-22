package model

type student struct {
	name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
}

// 通过工厂模式解决结构体名称小写导致外部导包无法访问的问题
func NewStudent(name string, age int, score float64) *student {
	return &student{
		name:  name,
		Age:   age,
		Score: score,
	}
}

// 通过工厂模式解决结构体字段名称因导包而导致无法访问的问题
func (stu *student) GetName() string {
	return stu.name
}

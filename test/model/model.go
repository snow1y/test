package model

type stu struct {
	Name string
}

func NewStu(n string) stu{
	return stu{n}
}

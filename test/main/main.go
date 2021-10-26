package main

import (
	"fmt"
	"test/model"
)
func main(){
	s := model.NewStu("hello")
	s.Name = "111"
	fmt.Println(s.Name)
}
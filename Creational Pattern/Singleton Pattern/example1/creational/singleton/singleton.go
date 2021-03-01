package singleton

import (
	"fmt"
)
// type Singleton interface {
// 	AddOne() int 
// }

type singleton struct {
	count int 
}
// định nghĩa 1 pointer trỏ đến struct của kiểu Singleton với giá trị nil và biến đó gọi là instance 
var instance *singleton

// fmt.Println("instance: ", instance)

func GetInstance() *singleton {
	if instance == nil {
		fmt.Println("instance1: ", instance)
		instance = new(singleton)
		fmt.Println("instance2: ", instance)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count 
}

 
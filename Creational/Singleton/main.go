package main

import "fmt"

// import "fmt"

/*
As an example of an object of which we must ensure that there is only one instance, we will
write a counter that holds the number of times it has been called during program execution.
*/

/*
When no counter has been created before, a new one is created with the value 0
If a counter has already been created, return this instance that holds the actual count
If we call the method AddOne, the count must be incremented by 1
*/

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

func main() {
	
	
	alo := GetInstance()
	fmt.Println("alo: ", alo)
}
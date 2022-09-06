package main

import "fmt"

type Human struct {
	Name string
	Age  int64
}

func (h Human) Pretty() string {
	return fmt.Sprintf("The human being with name %s and age %d", h.Name, h.Age)
}

type Action struct {
	Human // Human is the embedded struct (composition)
}

func main() {
	// instances of Action have Human fields and method
	a := Action{}
	a.Name = "Mark"
	a.Age = 24

	fmt.Println(a.Pretty())
}

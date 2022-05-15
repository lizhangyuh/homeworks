package main

import "fmt"

type T struct {
	Name string
}

func main() {
	v_new := new(T)
	v_new.Name = "jerry"

	v_1 := T{}
	v_1.Name = "jerry"

	fmt.Printf("%v, %T, %p\n", v_new, v_new, v_new)
	fmt.Printf("%v, %T, %p\n", &v_new, &v_new, &v_new)
}

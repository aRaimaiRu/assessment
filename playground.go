package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	var a int

	ptf(&a)
	paradicc(&a)
	fmt.Println(a)

}
func ptf(a *int) {

	// dereferencing
	*a = 748
}
func paradicc(arg ...any) {
	value, ok := arg[0].(*int)
	if !ok {
		fmt.Println("arg[0] is not an int")
	} else {
		*value = 743
	}
}

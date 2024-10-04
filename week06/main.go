package main

import (
	"fmt"
)

func main() {
	var b bool
	var s string
	var i int
	var f float64

	fmt.Println(f, b, s, i)
	fmt.Printf("%f%t%s%d", f, b, s, i)
	f = 2.7
	i = 3
	fmt.Print("\n\n", f > float64(i), "\n")
	//i := 13
	//var f float64 = 12.9
	//c1 := 'A'
	//c2 := '김'
	//fmt.Printf("value i : %d, value f : %f\n", i, f)
	//fmt.Printf("%d * %f = %d\n", i, f, i*int(f))

	//fmt.Println(c1, c2)
	//fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2))
	//fmt.Printf("value i: %d. value f : %f\n", i, f)
	//fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)
	//fmt.Println(reflect.TypeOf(i))
}

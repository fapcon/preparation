package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["s"] = 21
	fmt.Println(m)

	fmt.Println(Sum(5.55, 3.33))
	fmt.Println(Sum(-5.55, 3.33))
	fmt.Println(Sum(-5, 21))

}

type Summable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Sum[T Summable](a T, b T) T {
	return a + b
}

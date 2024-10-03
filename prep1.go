package main

import (
	"fmt"
	"sync"
)

//func changeP(p *int) {
//	v := 3
//	*p = v
//}
//
//func main() {
//	v := 5
//	p := &v
//	fmt.Println(*p)
//	changeP(p)
//	fmt.Println(*p)
//}

//func main() {
//	per := &Person{
//		Name: "Gena",
//	}
//	fmt.Println(per.Name)
//	changeName(per)
//	fmt.Println(per.Name)
//
//}
//
//type Person struct {
//	Name string
//}
//
//func changeName(per *Person) {
//	*per = Person{Name: "Olga"}
//}

//func main() {
//	var x []int
//	x = append(x, 0)  // {0}, 1, 1
//	x = append(x, 1)  // {0,1}, 2, 2
//	x = append(x, 2)  // {0,1,2} , 3, 4
//	y := append(x, 3) // {0,1,2,3}, 4, 4
//	z := append(x, 4) // {0,1,2,4}, 4, 4
//	fmt.Println(y, z) // {0,1,2,4}, {0,1,2,4} y и z ссылаются на одну и ту же область памяти
//}

//func main() {
//	var x []int
//	x = append(x, 0)
//	x = append(x, 1)
//	x = append(x, 2)  //{0,1,2} ,3 ,4
//	x = append(x, 3)  //{0,1,2,3}, 4, 4
//	y := append(x, 4) //{0,1,2,3,4}, 5, 8
//	z := append(x, 5) //{0,1,2,3,5}, 5, 8
//	fmt.Println(y, z) // {0,1,2,3,4}, {0,1,2,3,5}
//}

//func main() {
//	var wg sync.WaitGroup
//	c := make(chan string, 3)
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//			c <- fmt.Sprintf("g %s", strconv.Itoa(i))
//		}(i)
//	}
//	go func() {
//		wg.Wait()
//		close(c)
//	}()
//	for v := range c {
//		fmt.Println(v)
//		//select {
//		//case v, ok := <-c:
//		//	if !ok {
//		//		return
//		//	}
//		//	fmt.Println(v)
//		//}
//	}
//}

//func main() {
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go run(&wg)
//	wg.Wait()
//}
//
//func run(wg *sync.WaitGroup) {
//	defer wg.Wait()
//	ch := make(chan int)
//	for i := 0; i < 3; i++ {
//		go func(idx int) {
//			ch <- idx + 1
//		}(i)
//	}
//	for {
//		select {
//		case v, ok := <-ch:
//			if !ok {
//				return
//			}
//			fmt.Println(v)
//		}
//	}
//}

//func main() {
//	c := make(C, 1)
//	c <- c
//	for i := 0; i < 1000; i++ {
//		select {
//		case <-c:
//		case <-c:
//			c <- c
//		default:
//			fmt.Println(i)
//			return
//		}
//	}
//}
//
//type C chan C

//func main() {
//	runtime.GOMAXPROCS(1)
//	ch := 0
//	go func() {
//		ch = 1
//	}()
//	for ch == 0 {
//	}
//	fmt.Println("finish")
//}

func main() {
	a := make(chan int)
	b := make(chan int)
	d := make(chan int)
	i := make(chan int)
	go func() {
		a <- 1
		a <- 101
		a <- 201
		a <- 301
		b <- 2
		d <- 3
		i <- 5
	}()
	c := merge(a, b, d, i)
	for {
		select {
		case v, ok := <-c:
			if !ok {
				return
			}
			fmt.Println(v)
		}
	}
}

func merge(a ...chan int) chan int {
	res := make(chan int)
	var wg sync.WaitGroup
	go func() {
		for _, ch := range a {
			wg.Add(1)
			go func() {
				for v := range ch {
					res <- v
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(res)
	}()
	return res
}

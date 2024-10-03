package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//testSlices1()
	//testSlices2()
	//testSlices3()
	//testSlices4()
	//testData1()
	//testData2()
	testGoroutines1()
}

func testSlices1() {
	a := []string{"a", "b", "c"}
	b := a[1:2]
	b[0] = "q"
	fmt.Printf("%s\n", a)
}

func testSlices2() {
	a := []byte{'a', 'b', 'c'}
	b := append(a[1:2], 'h')
	b[0] = 'z'
	fmt.Printf("%s\n", a)
}

func testSlices3() {
	a := []byte{'a', 'b', 'c'}
	b := append(a[1:2], 'd', 'x')
	b[0] = 'z'
	fmt.Printf("%s\n", a)
}

func testSlices4() {
	a := []byte{'a', 'b', 'c'}
	b := string(a)
	a[0] = 'z'
	fmt.Printf("%s\n", b)
}

type MyData struct {
	One int    `json:"one",qwe:"123"`
	two string `json:"two"`
}

func testData1() {
	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in)
	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded))
	var out MyData
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out)
}

func testData2() {
	a := []int{1, 2, 3, 4}
	result := make([]*int, len(a))
	for i, v := range a {
		result[i] = &v
	}
	for _, u := range result {
		fmt.Printf("%d ", *u)
	}
}

func testGoroutines1() {
	//var ch chan int
	ch := make(chan int)
	for i := 0; i < 30000000000; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
		//fmt.Println("result:", <-ch)
	}
	fmt.Println("result:", <-ch)
	time.Sleep(2 * time.Second)
}

func testGoroutines2() {
	ch := make(chan string)
	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
		}
	}()
	ch <- "cmd.1"
	ch <- "cmd.2"
}

func testGoroutines3() {
	var num int
	wg := &sync.WaitGroup{}
	for i := 0; i < 10000; i++ {

		defer wg.Done()
		num = i
	}
	//wg.Wait()
	fmt.Printf("NUM is %d", num)
}

func testGoroutines4() {
	dataMap := make(map[string]int)
	for i := 0; i < 10000; i++ {
		go func(d map[string]int, num int) {
			d[fmt.Sprintf("%d", num)] = num
		}(dataMap, i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(len(dataMap))
}

func testGoroutines5() {
	runtime.GOMAXPROCS(1)

	x := 0
	go func(p *int) {
		for i := 1; i <= 20000000000; i++ {
			*p = i
			runtime.Gosched()
		}
	}(&x)

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("x = %d.\n", x)
}

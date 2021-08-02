package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 47
		ch <- 21
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()

}

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	fmt.Printf("%v\n", runtime.GOMAXPROCS(-1))
// }

// import (
// 	"fmt"
// )

// func main() {
// 	defer fmt.Printf("defer1\n")
// 	defer fmt.Printf("defer2\n")

// 	var i interface{} = 1.0
// 	switch i.(type) {
// 	case int:
// 		fmt.Printf("int\n")
// 	case float64:
// 		fmt.Printf("float64\n")
// 	default:
// 		fmt.Printf("default\n")
// 	}

// }

// func main() {
// 	a := make([]int, 3)
// 	fmt.Printf("%v, %T\n", a, a)
// 	// grades := []int{97, 66, 52}
// 	// grades1 := grades
// 	// grades1[0] = 100
// 	// fmt.Printf("%v, %T\n", grades, grades)
// 	// fmt.Printf("%v, %T\n", grades1, grades1)

// 	// grades := [...]int{97, 66, 52}
// 	// grades1 := &grades
// 	// grades1[0] = 100
// 	// fmt.Printf("%v, %T\n", grades, grades)
// 	// fmt.Printf("%v, %T\n", grades1, grades1)
// }

// import (
// 	"fmt"
// 	// "strconv"
// )

// var (
// 	name  string  = "666"
// 	Int   int     = 666
// 	Float float32 = 666.6
// )

// const (
// 	a = iota
// 	b
// 	c
// )

// const (
// 	a2 = iota
// )

// func main() {
// 	// i := 96

// 	// s := strconv.Itoa(i)
// 	// fmt.Printf("%v, %T", s, s)

// 	// b := true
// 	// fmt.Printf("%v, %T", b, b)

// 	// var n complex64 = 1 + 2i
// 	// var n complex64 = complex(1, 2)
// 	// fmt.Printf("%v, %T\n", n, n)
// 	// fmt.Printf("%v, %T\n", real(n), real(n))
// 	// fmt.Printf("%v, %T\n", imag(n), imag(n))

// 	// s := "string"
// 	// b := []byte(s)
// 	// fmt.Printf("%v, %T\n", b, b)

// 	fmt.Printf("%v, %T\n", a, a)
// 	fmt.Printf("%v, %T\n", b, b)
// 	fmt.Printf("%v, %T\n", c, c)

// 	fmt.Printf("%v, %T\n", a2, a2)
// }

package main

// func main() {
// 	scanner := bufio.NewScanner(strings.NewReader(`one
// two
// three
// four
// `))
// 	var (
// 		text string
// 		n    int
// 	)
// 	for scanner.Scan() {
// 		n++
// 		text += fmt.Sprintf("%d. %s\n", n, scanner.Text())
// 	}
// 	fmt.Print(text)
// }

// type Orange struct {
// 	Quantity int
// }

// func (o *Orange) Increase(n int) {
// 	o.Quantity += n
// }

// func (o *Orange) Decrease(n int) {
// 	o.Quantity -= n
// }

// func (o *Orange) String() string {
// 	return fmt.Sprintf("%v", o.Quantity)
// }

// func main() {
// 	var orange Orange
// 	orange.Increase(10)
// 	orange.Decrease(5)
// 	fmt.Println(orange)
// }

// func main() {
// 	runtime.GOMAXPROCS(1)

// 	done := false

// 	go func() {
// 		done = true
// 	}()

// 	for !done {
// 	}
// 	fmt.Println("finished")
// }
// func main() {
// var counter int
// for i := 0; i < 1000; i++ {
// 	go func() {
// 		counter++
// 	}()
// }
// 	v := 5
// 	p := &v
// 	println(*p)

// 	changePointer(p)
// 	println(*p)
// }

// func changePointer(p *int) {
// 	v := 3
// 	p = &v
// }
// func worker() chan int {
// 	ch := make(chan int)

// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		ch <- 42
// 	}()

// 	return ch
// }

// func main() {
// 	timeStart := time.Now()

// 	go worker()
// 	go worker()

// 	println(int(time.Since(timeStart).Seconds())) // 3 or 6 ?
// }

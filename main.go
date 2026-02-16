package main

import (
	"fmt"
	"time"
)

type Queue struct {
	Id   int
	Name string
	Wait int
}

func main() {
	// fmt.Println("hahahaah")
	// var wg sync.WaitGroup

	queue := []Queue{
		{
			Id:   1,
			Name: "test",
			Wait: 9,
		},
		{
			Id:   2,
			Name: "Hello",
			Wait: 3,
		},
		{
			Id:   3,
			Name: "World",
			Wait: 5,
		},
	}

	// times := time.Second
	// fmt.Println(times)
	// fmt.Println(time.Sleep(5))

	start := time.Now()

	fmt.Println("bersiaplah")

	for x := range queue {
		time.Sleep(time.Duration(queue[x].Wait) * time.Second)

		duration := time.Since(start)
		fmt.Print("\n\n\nPesanan: ", queue[x].Name, " selesai dalam ", duration.Seconds(), "\n")

		// fmt.Println("time elapsed in minutes:", duration.Minutes())
		// fmt.Println("time elapsed in hours:", duration.Hours())

	}

	// for x := range queue {
	// 	value := make(chan []string)
	// 	go func() {
	// 		defer wg.Done()
	// 		rawResult := queue[x].Name
	// 		// value = <-rawResult
	// 	}()
	// }

	// for x := range queue {
	// 	value := make(chan Queue)
	// 	go func() {
	// 		defer wg.Done()
	// 		fmt.Println(queue[x])
	// 		value <- queue[x]
	// 	}()
	// 	result := <-value
	// 	fmt.Println(result)
	// }
	// wg.Wait()

	// go func() {
	// 	fmt.Println(12)
	// }()

}

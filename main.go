package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	Id   int
	Name string
	Wait int
}

func main() {
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

	fmt.Println("Pemesanan Ayam Goyek")

	// start := time.Now()
	// for x := range queue {
	// 	time.Sleep(time.Duration(queue[x].Wait) * time.Second)
	// 	duration := time.Since(start)
	// 	fmt.Print("\n\n\nPesanan: ", queue[x].Name, " selesai dalam ", queue[x].Wait, "\n\n")
	// 	fmt.Print("\n\n semuanya selesai dalam waktu ", duration.Seconds())
	// }

	var wg sync.WaitGroup

	wg.Add(len(queue))
	cen := make(chan string, len(queue))
	go func() {
		for x := range queue {
			defer wg.Done()
			cen <- queue[x].Name
		}
	}()
	wg.Wait()
	for x := range len(cen) {
		c := <-cen
		time.Sleep(time.Duration(queue[x].Wait) * time.Second)
		fmt.Print("\nPesanan: ", c, " selesai dalam ", queue[x].Wait, " Detik \n\n")

	}
}

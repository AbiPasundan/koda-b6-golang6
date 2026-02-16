package main

import "fmt"

type Queue struct {
	Id   int
	Name string
	Wait int
}

func main() {
	// fmt.Println("hahahaah")

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

	fmt.Println(queue[0].Id)

	for x := range queue {
		go func() {
			fmt.Println(12)
		}()
	}

	// go func() {
	// 	fmt.Println(12)
	// }()

}

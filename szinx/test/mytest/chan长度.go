package main

import "fmt"

func main()  {
	nums := make(chan int,10)
	fmt.Println(cap(nums))
}


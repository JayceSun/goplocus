package main

import "fmt"


const (
	UNSET int = iota - 1
	OFFLINE
	SERVER
	CLIENT
)

func pump(ch chan int) {
	for i := 0;i < 10 ; i++ { ch <- i } //往管道里写值
}

func suck(ch chan int) {
	for { fmt.Println(<-ch) } //这里会等着直到有值从管道里面出来
}

func main() {

	ch := make(chan int) //创建一个只能传递整型的管道

	go pump(ch) //异步执行pump

	go suck(ch) //异步执行suck
}

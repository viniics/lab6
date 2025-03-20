package main

import ("fmt"
		"math/rand"
		"time"
	)

func produce(ch chan int){
	for{
		n := rand.Intn(10)
		fmt.Println(n, "Foi produzido")
		ch <- n
		time.Sleep(200*time.Millisecond)
	}
}

func consume(ch chan int){
	for{
	n:= <- ch
	if n %2==0{fmt.Println(n, "Foi consumido")}
}
}

func main(){
	ch := make(chan int)
	join:= make(chan int)
	go produce(ch)
	go consume(ch)
	<- join
}
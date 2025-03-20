package main

import ("fmt"
		"math/rand"
		"time"
	)

func produce(ch chan int, max int){
	for range max{
		n := rand.Intn(10)
		fmt.Println(n, "Foi produzido")
		ch <- n
		time.Sleep(200*time.Millisecond)
	}
	close(ch)
}

func consume(ch chan int, join chan int){
	for {
		n, isClosed := <- ch
		if !isClosed {
			join <- 4
			break
		}
		if n %2==0{fmt.Println(n, "Foi consumido")}
	}
}

func main(){
	ch := make(chan int)
	join:= make(chan int)
	go produce(ch,rand.Intn(10))
	go produce(ch,rand.Intn(10))
	go consume(ch, join)
	<- join
}
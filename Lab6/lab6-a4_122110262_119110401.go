package main

import ("fmt"
		"math/rand"
		"time"
	)

func produce(ch chan<- int, max int){
	for range max{
		n := rand.Intn(10)
		fmt.Println(n, "Foi produzido")
		ch <- n
		time.Sleep(200*time.Millisecond)
	}
	close(ch)
}

func consume(ch <-chan int, join chan int){
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
	ch := make(chan int,10)
	join:= make(chan int,1)
	go produce(ch,rand.Intn(100))
	go produce(ch,rand.Intn(100))
	go consume(ch, join)
	<- join
}
package main

import ("fmt"
		"math/rand"
		"time"
	)

func gateway(ngo int, wait_n int,ch chan int, alreadyFinished chan int) int {
	soma := 0

	for range ngo{
		go request(ch,alreadyFinished)
	}

	for{
		n:= len(alreadyFinished)
		if(n == ngo){
			break
		}
	}

	for valor:= range(ch){
		soma+= valor
	}

	return soma
}

func request(ch chan<- int, alreadyFinished chan <- int){
	n:= rand.Intn(1)
	ch <- n
	fmt.Println("Num Sorteado", n)
	time.Sleep(time.Duration(n)*time.Second)
	alreadyFinished <- 1
}

func main(){
	ch:= make(chan int)
	alreadyFinished := make(chan int)
	n:= gateway(rand.Intn(10),1,ch,alreadyFinished)
	fmt.Println(n)
}
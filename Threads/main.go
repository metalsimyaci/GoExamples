package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Başladı")
	runtime.GOMAXPROCS(1)
	go xFunc()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Bitti")
}
func xFunc()  {
	for l := byte('a'); l <= byte('z'); l++  {
		go fmt.Println(string(l))
	}
}

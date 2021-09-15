package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganho!", i)
	} else {
		fmt.Println("Perdeu!", i)
	}
}

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

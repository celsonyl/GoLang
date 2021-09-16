package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(2))
	fmt.Println(tipo("oxi"))
	fmt.Println(tipo('i'))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Float"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não sei"
	}
}

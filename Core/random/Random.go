package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("рандомное число")
	rand.Seed(time.Now().UnixNano()) // чтобы число реально генерировалось случайно и не повторялось
	randomValue := rand.Intn(100)    // само генерация числа где 5 это максимальное число для выборки 1 это минимальное число, то есть от 1 до 4
	fmt.Println(randomValue, "Rand")

}

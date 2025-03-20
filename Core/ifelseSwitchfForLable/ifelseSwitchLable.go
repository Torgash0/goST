package main

import "fmt"

func main() {
	fmt.Println("if/switch")
	/*
			if "условие"{
				и здесь мы можем наисать continue если нам надо пропустить
				}
			 switch value{
		case 1: fallthough  это мы пишем если нам надо дальше смотреть по кейсам , то есть мы зашли в case 1, выполнили его и после этого мы пойдем дальше , все благодаря fallthrough
		case 2:
		default: // это условие , если ничего не подошло из case
	*/
	/*
		for { // это бесконенчный цикл
		//do somfing
		}
	*/
	num := 5
	ages := map[string]int{ //создание мапы
		"Andru": 20,
		"Mike":  19,
	}
	var s string = "слово"

	for num == 1 { //цикл с условием
		//do somfing
	}

	for i := 0; i < 5; i++ { //цикл с парметром
		//do somfing
	}

	sl := []int64{9, 2, 7, 4, 5}
	fmt.Println("работа с slice")

	for key, value := range sl { //в данном случе range интерируется по нашему слайсу sl так, что мы модем достать из него и ключ и значение
		fmt.Printf("key:%v, val: %v", key, value)
	}
	for _, value := range sl { //это если нам нужено игнорировать к примеру ключ, то мы пишем вместо key нижние подчеркивание
		fmt.Printf("val: %v", value)
	}
	fmt.Println("работа с MAP")
	for key, value := range ages { //это исопльзование for и range для MAP
		fmt.Println(key)
		fmt.Println(value)
	}
	fmt.Println("работа с string не правильная ")
	for i := 0; i < len(s); i++ { // так не надо ходить по строке, тут мы интерируемся по байтно
		fmt.Println(s[i])
		fmt.Printf("%T", s[i])
	}
	fmt.Println("работа с string правильная")
	for key, value := range s { //вот так НАДО, тут мы итерируемся по символьно
		fmt.Println(key)   // индекс нашего символа
		fmt.Println(value) //байтовое представлнеие нашего символа
		//fmt.Printf("%T", s[i])
	}
	fmt.Println("////////////////////")

	fmt.Println("пример с метками и continue ")

	//работа с continue и break
	//разница между continue и break это то, что continue пропускает интерацию по циклу, а break завершает работу с циклом
Lable1: // работа с метками и continue
	for i := 0; i < 20; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, j)
			if i >= 10 {
				fmt.Println("continue Lable1")
				continue Lable1
			}
		}
	}
	fmt.Println("пример с метками и break ")
Lable2: // работа с метками и break
	for i := 0; i < 20; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, j)
			if i >= 10 {
				fmt.Println("continue Lable2")
				break Lable2
			}
		}
	}

}

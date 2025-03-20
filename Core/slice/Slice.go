package main

import "fmt"

func main() {
	fmt.Println("SLICE")
	var nums []int = []int{1, 2, 3, 4, 4, 5} // обьявление среза он же слай, размер определяется динамически
	fmt.Println("доступ к числу по индексу", nums[2])
	numbers := nums[2:4]
	numbersLen := len(numbers) //узнаем длинну выборки
	numbersCap := cap(numbers) //узнаем размер капасити ,емкость нашего среза, благодря этой характеристики мы можем удлинять нащ срез/слайс на это число
	list := make([]int, 0, 5)  //создание нового среза, первое число говорит о длинне среза (len), втрое число гооврит о емкости среза(cap)
	list = append(list, 6)     //добавление нового элемента в срез, если емкость слайса закончилась , то он пересоздасться и туда все эелементы скопируються , а размер стнет в 2 раза больше
	fmt.Println(list)
	list = append(list, nums...) //добавление одного слайса в другой
	fmt.Println(nums, numbers, numbersLen, numbersCap)
	copyslice := make([]int, len(list), len(list)) //копировнаие слайса list в слайс copyslice
	copy(copyslice, list)                          //копировнаие слайса list в слайс copyslice
	numsL := len(nums)
	fmt.Println(list, numsL)

	//Проход по слайсу и проверка на наличе в слайсе повторяющихся элементо
	fmt.Println("Проход по слайсу и проверка на наличе в слайсе повторяющихся элементов")
	counts := make(map[int]int)

	for _, num := range nums {
		fmt.Println(counts, "////")
		counts[num]++
		fmt.Println(counts)
		if counts[num] >= 2 {
			fmt.Println("true") // Если элемент встречается более одного раза, возвращаем true
			break
		}
	}
	findElement(nums, 1)

}
func findElement(nums []int, s int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == s {
			fmt.Println("есть такой элемент")
		}

	}
}

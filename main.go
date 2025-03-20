package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
	"unicode/utf8"
)

// это для структур

type Avatar struct {
	URL  string
	Size string
}
type Person struct {
	Name string
	Age  int
}
type Client struct {
	ID  int64
	IMG Avatar
	Person
}

// для интерфейсов
type Sender interface {
	Send(msg string) error
}

type Email struct {
	Adress string
}

type Phone struct {
	Number  int
	Balance int
}

func (e *Email) Send(msg string) error {
	fmt.Printf("Sending email to %s\n", msg, e.Adress)
	return nil
}
func (p *Phone) Send(msg string) error {
	fmt.Printf("Sending phone to %s\n", msg, p.Number)
	return nil
}

//func Notify(s Sender) {
//	err := s.Send("Notify messege"
//	if err != nil {
//		fmt.Printf("Error: %s\n", err)
//		return
//	}
//	//email, ok := i.(Email) // это и есть утверждление типов
//	switch s.(type) { //такая штука назвается type switch (это только для утверждения типов)
//	case *Email: // тут утверждение типов
//		fmt.Println("send email")
//	case *Phone:
//		fmt.Println("send phone")
//
//	}
//	phone, ok := s.(*Phone) // это утверждение типа, такая штука нам нудеа только в функциях по тиму этой , если при какой то структуре нам нужно работтаь с переменной этой структуры
//	if ok {
//		fmt.Println(phone.Balance)
//	}
//	fmt.Printf("sucsess")
//}

// это для пустого интефеса
func Notify(i interface{}) {
	switch i.(type) { //такая штука назвается type switch (это только для утверждения типов)
	case int:
		fmt.Println("число не поддреживается")
	}
	s, ok := i.(Sender) //проверка на принадледность к интерфесу
	if !ok {
		fmt.Println("Error1")
		return
	}
	err := s.Send("сообщение пустого интерфейса") // предалось ли мообщениее
	if err != nil {
		fmt.Printf("Error2")
		return
	}
	fmt.Printf("sucsess")
}

func main() {

	//ПЕРЕМЕННЫЕ
	var s string = "слово"
	var i int = 5                        //обьявление переменной численной
	f := 6                               //можно и так обьявлять перпемнные
	var intPointer *int = &i             //обьявление указателя
	x := &i                              //присвоение  указателя
	*x = 10                              // изменение числа по указателю
	fmt.Println(i, f, i, *x, intPointer) //вывод в косоль
	var String string = "Строка"         //обьявление строки ( строку нельзя изменять, при изменении строка создаеться заново
	Stri1 := String[2:4]                 //выборка между букв
	Stri2 := String[:4]                  //выборка до 4 буквы
	Stri3 := String[4:]                  //выборка после 4 буквы включительно

	st := len(String)                       //возвращает длинну строки в байтах , кодировка либо 8 байт либо 32
	strln := utf8.RuneCountInString(String) // а это уже длинна строки
	fmt.Println(String, Stri1, Stri2, Stri3, st, strln)

	fmt.Println("////////////////////")

	fmt.Println("ARRAY")
	//МАССИВЫ
	var mas [5]int = [5]int{1, 2, 3, 4, 5}                        //обьявление массива
	var mass = [5]int{1, 2, 3, 4, 5}                              //так же мы можем обьявить его и так
	j := &mass                                                    // указатель на массив
	masss := [5]int{1, 2, 3, 4, 5}                                //так же мы можем обьявить его и так
	massss := [...]string{"один", "два", "три", "четыре", "пять"} //четырмя точки мы задем длинну , если не знаем ее точно
	m := len(massss)                                              //функция возвращает длинну массива в байтах, одировка либо 8 байт либо 32
	fmt.Println(mas, mass, masss, massss, m, *j)

	fmt.Println("////////////////////")

	//SLICE
	fmt.Println("SLICE")
	var nums []int = []int{1, 2, 4, 4, 5} // обьявление среза он же слай, размер определяется динамически
	var slString = []string{"a", "s", "d", "f", "g", "h", "j", "k", "l"}
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

	fmt.Println("////////////////////")

	fmt.Println("MAP")
	//MAP
	var m1 map[int32]int     //создание мапы
	var m2 map[string]string //создание мапы
	m3 := make(map[int]int)  //создание мапы
	ages := map[string]int{  //создание мапы
		"Andru": 20,
		"Mike":  19,
	}
	age := ages["Andru"] // так мы получим значение по ключу
	ages["Vera"] = 23    // добавление нового значения
	agesln := len(ages)
	age, ageExists := ages["Sergey"] //проверка есть ли в мапе этот ключ
	fmt.Println(age, ageExists)
	fmt.Println(m1, m2, m3, ages, age, ages, agesln)
	delete(ages, "Andre") //удаление элемента из мапы

	fmt.Println("////////////////////")

	//STRUCT
	fmt.Println("STRUCT")
	type point struct { //обьвление структуры и полей, которые будут входить в нее
		X, Y int
		//Y int
	}
	p := point{ //создание структуры
		X: 1,
		Y: 2,
	}
	q := point{3, 4}           //создание структуры(аналог)
	p.X = 5                    //перепресваивание полей
	qPointer := &q             //получение полкй структруы по ее указателю
	fmt.Println((*qPointer).X) //можео так выводить значение указателя
	fmt.Println(qPointer.X)    //НО ЛУЧШЕ ТАК ЭТО ДЕЛАТЬ
	fmt.Println(p, p.X, p.Y, q, q.X, q.Y)

	w := new(int64) //выделения памяти и обьявление указателя на этот участок памяти
	_ = w
	client := Client{}

	fmt.Printf("%#v\n", client)
	updateAvatar(&client) //предеча по ссылке структуры клиент
	fmt.Printf("%#v\n", client)
	updateClient(&client)

	fmt.Println("////////////////////")

	//Method
	fmt.Println("методы")
	clientM := &Client{
		ID: 1,
		IMG: Avatar{
			URL:  "some_url",
			Size: "10",
		},
		Person: Person{
			Name: "Sergey",
			Age:  20,
		},
	}
	//clientM.person = Person{"Sergey", 20}
	//clientM.Name = "Sergey"
	//clientM.Age = 20
	fmt.Println(clientM.Age, client.IMG.URL)
	personNew := newPerson("setgey", 10)
	fmt.Println(personNew.Name, personNew.Age)
	fmt.Println(clientM.HasAvatar())
	clientM.UpdateAvatar() //для обновления нужно работать с сылкой на структуру, а не с самой
	clientM.UpdateId(10)
	fmt.Println(clientM)

	fmt.Println("////////////////////")

	// FUNC
	fmt.Println("FUNC")
	myString := addPrefix("my_string")
	fmt.Println(myString)
	myString, err := addPrefixWithErr("my_string")
	fmt.Println(err)
	fmt.Println(myString)
	var multiplier func(x, y int) int                //использование функции как занчения первый варинат
	multiplier = func(x, y int) int { return x * y } //использование функции как занчения первый варинат
	divider := func(x, y int) int { return x / y }   //использование функции как занчения второй  варинат
	fmt.Println(multiplier(5, 5), divider(10, 2))
	fmt.Println(calculate(5, 2, multiplier))
	dividerBy2 := createDivider(2) //возврат функции из функции
	fmt.Println(dividerBy2(10))
	slTestFunc := addSlice([]string{})
	slTestFuncString := addSlice(slString)
	fmt.Println(slTestFunc, slTestFuncString)
	//DEFER ( работда с FUNC)
	defer end() //это функция которая будет выполнена после заверщение основной функции ( проще говоря в конце), ее используют для закрытия конектов и закрытия файлов и для обработки паник
	num := 5
	defer func(x int) { //это анонимная функция с дефером (обратить внимание на (num) в конце
		fmt.Println(x)
	}(num)
	num = 10

	fmt.Println("////////////////////")

	//if and swich case
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

	fmt.Println("////////////////////")

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
	fmt.Println("////////////////////")

	//Random
	fmt.Println("рандомное число")
	rand.Seed(time.Now().UnixNano()) // чтобы число реально генерировалось случайно и не повторялось
	randomValue := rand.Intn(100)    // само генерация числа где 5 это максимальное число для выборки 1 это минимальное число, то есть от 1 до 4
	fmt.Println(randomValue, "Rand")

	fmt.Println("////////////////////")

	//interface
	fmt.Println("INTERFACE")

	emeil := &Email{"dev@gmail.com"}           //мы не моожем передавать в интерфес структуру, только узказатель
	phone := &Phone{Number: 7777, Balance: 10} //мы не моожем передавать в интерфес структуру, только узказатель
	slInterface := [5]int64{1, 2, 3, 4, 5}

	Notify(slInterface)
	Notify(emeil)
	Notify(phone)
	fmt.Println("")
	fmt.Println("////////////////////")

	//работа с файлами

	//первый способ нахождения уникальных элементов
	//in := bufio.NewScanner(os.Stdin)
	//alreadySeen := make(map[string]bool)
	//for in.Scan() {
	//	txt := in.Text()
	//	if _, found := alreadySeen[txt]; found {
	//		continue
	//	}
	//	alreadySeen[txt] = true
	//	fmt.Println(txt)
	//}

	//	//второй способ нахождения уникальных элементов
	//	in := bufio.NewScanner(os.Stdin)
	//	var prev string
	//	for in.Scan() {
	//		txt := in.Text()
	//		if txt == prev {
	//			continue
	//		}
	//		if txt < prev {
	//			panic("file not sorted")
	//		}
	//		prev = txt
	//		fmt.Println(txt)
	//	}
	//
	errr := uniq(os.Stdin, os.Stdout)
	if errr != nil {
		panic(errr.Error())
	}

	fmt.Println("////////////////////")
}

// для способ нахождения уникальных элементов и теста
func uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prev string
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			return fmt.Errorf("file not sorted")
		}
		prev = txt
		fmt.Println(output, txt)
	}
	return nil
}

// FUNC
func addSlice(sl []string) []string {
	fmt.Println(append(sl, "slice"))
	return sl
}
func addPrefix(origin string) string { // там где написанно (origin string) это то , что мы передаем , а где просто string это то,что мы возвращаем
	return "prefix_" + origin
}
func addPrefixWithErr(origin string) (string, error) { // там где написанно (origin string) это то , что мы передаем , а где (string, error) это то,что мы возвращаем
	return "prefix_with_err_" + origin, nil
}
func addPrefixWithLen(origin string) (res string, lenght int) { //если мы явно обьяляем о возвращаймых перменных ,
	//  то мы можем не указывать ничего в ретерн, так как при таком явном указании переменным присваиваться их стандарнтые значения в случае строки это "" ,
	// а в случае инта это 0 (для нашего понимания можем указать)
	res = "prefix_" + origin
	lenght = len(origin)
	return
}
func end() { //это для примера с defer
	fmt.Println("end")
}
func calculate(x, y int, action func(int, int) int) int { // когда в функцию помещаем другую функцию
	return action(x, y)
}
func createDivider(divider int) func(x int) int { //возврат функции из функции
	dividerFunc := func(x int) int {
		return x / divider
	}
	return dividerFunc
}

// для структур
func updateAvatar(client *Client) {
	client.IMG.URL = "update_url"
}
func updateClient(client *Client) {
	client.Name = "Vera"
}
func newPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// Method
func (c Client) HasAvatar() bool { // для проверки можно работать с самой структурой
	if c.IMG.URL != "" {
		return true
	}
	return false
}

func (c *Client) UpdateAvatar() { //для обновления нужно работать с сылкой на структуру, а не с самой
	c.IMG.URL = "update_url"
}
func (c *Client) UpdateId(id int64) {
	c.ID = id
}

//Стек:
//Используйте стек для локальных переменных и временных данных, которые не нужны после завершения функции.
//Стек подходит для небольших объектов и данных, которые не требуют долгосрочного хранения.
//Основные концепции работы стека в Go:
//Вызов функции:
//Когда функция вызывается, информация о вызове (например, адрес возврата, параметры функции и локальные переменные) сохраняется в стеке.
//Это позволяет функции выполнять свои операции и затем вернуться к вызывающей функции.
//Локальные переменные:
//Локальные переменные функции хранятся в стеке. Они существуют только в течение выполнения функции и автоматически освобождаются, когда функция завершает выполнение.
//Адрес возврата:
//Адрес возврата — это адрес инструкции, к которой нужно вернуться после завершения выполнения функции. Этот адрес также сохраняется в стеке.
//Управление памятью:
//Стек управляется автоматически и освобождается, когда функция завершает выполнение. Это позволяет избежать утечек памяти и других проблем, связанных с управлением памятью.

//Куча:
//Используйте кучу для объектов, которые должны существовать дольше, чем время выполнения функции.
//Куча подходит для больших объектов и данных, которые требуют долгосрочного хранения.
//В Go куча (heap) — это область памяти, используемая для динамического выделения памяти во время выполнения программы.
//Куча управляется сборщиком мусора (garbage collector), который автоматически освобождает память, когда она больше не используется.
//Используйте кучу для объектов, на которые ссылаются указатели, чтобы избежать копирования больших объектов.
//Основные концепции работы с кучей в Go:
//Выделение памяти:
//В Go память в куче выделяется с помощью встроенных функций new и make, а также с помощью оператора &.
//Сборка мусора:
//Go имеет встроенный сборщик мусора, который автоматически освобождает память, когда она больше не используется. Это
//позволяет избежать утечек памяти и других проблем, связанных с управлением памятью.
//Указатели:
//В Go указатели используются для работы с памятью в куче. Указатели позволяют передавать ссылки на объекты, что позволяет
//избежать копирования больших объектов и улучшить производительность.

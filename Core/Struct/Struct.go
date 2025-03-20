package main

import "fmt"

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
	IMG Ava	tar
	Person
}

func main() {
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
	_ = client

	fmt.Printf("%#v\n", client)
	updateAvatar(&client) //предеча по ссылке структуры клиент
	fmt.Printf("%#v\n", client)
	updateClient(&client)

	//METHOD
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
	//DEFER ( работда с FUNC)
	defer end() //это функция которая будет выполнена после заверщение основной функции ( проще говоря в конце), ее используют для закрытия конектов и закрытия файлов и для обработки паник
	num := 5
	defer func(x int) { //это анонимная функция с дефером (обратить внимание на (num) в конце
		fmt.Println(x)
	}(num)
	num = 10

}

// FUNC
func addPrefix(origin string) string { // там где написанно (origin string) это то , что мы передаем , а где просто string это то,что мы возвращаем
	return "prefix_" + origin
}
func addPrefixWithErr(origin string) (string, error) { // там где написанно (origin string) это то , что мы передаем , а где (string, error) это то,что мы возвращаем
	return "prefix_with_err_" + origin, nil
}
func addPrefixWithLen(origin string) (res string, lenght int) { //если мы явно обьяляем о возвращаймых перменных ,  то мы можем не указывать ничего в ретерн, так как при таком явном указании переменным присваиваться их стандарнтые значения в случае строки это "",а в случае инта это 0 (для нашего понимания можем указать)
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
func (c *Client) HasAvatar() bool { // для проверки можно работать с самой структурой
	if c.IMG.URL != "" {
		return true
	}
	return false
}

func (c *Client) UpdateAvatar() { //для обновления нужно работать с сылкой на структуру, а не с самой
	c.IMG.URL = "update_url"
}

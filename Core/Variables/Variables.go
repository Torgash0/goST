package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const cnst string = "name"

func main() {
	//ПЕРЕМЕННЫЕ

	fmt.Println("переменные")
	var str string = "слово"
	var i int = 5            //обьявление переменной численной
	f := 6                   //можно и так обьявлять перпемнные
	var intPointer *int = &i //обьявление указателя
	x := &i                  //присвоение  указателя
	*x = 10                  // изменение числа по указателю
	a := 3
	b := 4
	a, b = b, a                                           //поменять местами переменныые
	fmt.Println(str, i, f, i, *x, intPointer, a, b, cnst) //вывод в косоль
	var String string = "Строка"                          //обьявление строки ( строку нельзя изменять, при изменении строка создаеться заново
	Stri1 := String[2:4]                                  //выборка между букв
	Stri2 := String[:4]                                   //выборка до 4 буквы
	Stri3 := String[4:]                                   //выборка после 4 буквы включительно

	st := len(String)                       //возвращает длинну строки в байтах , кодировка либо 8 байт либо 32
	strln := utf8.RuneCountInString(String) // а это уже длинна строки
	fmt.Println(String, Stri1, Stri2, Stri3, st, strln)

	// 1. Создание строк
	s := "Hello, World!"
	fmt.Println("Original string:", s)

	// 2. Доступ к символам строки
	firstChar := s[0] // Байт, а не символ
	fmt.Println("First byte:", firstChar)

	// 3. Срезы строк
	substring := s[0:5] // "Hello"
	fmt.Println("Substring:", substring)

	// 4. Преобразование строк в руны
	runes := []rune(s)
	firstRune := runes[0] // Символ 'H'
	fmt.Println("First rune:", firstRune)

	// 5. Преобразование строк в байты
	bytes := []byte(s)
	firstByte := bytes[0] // Байт 'H'
	fmt.Println("First byte:", firstByte)

	// 6. Конкатенация строк
	s1 := "Hello"
	s2 := "World"
	result := s1 + ", " + s2 // "Hello, World"
	fmt.Println("Concatenated string:", result)

	// 7. Функции пакета `strings`
	contains := strings.Contains(s, "World") // true
	fmt.Println("Contains 'World':", contains)

	count := strings.Count(s, "l") // 3
	fmt.Println("Count of 'l':", count)

	hasPrefix := strings.HasPrefix(s, "Hello") // true
	fmt.Println("Has prefix 'Hello':", hasPrefix)

	hasSuffix := strings.HasSuffix(s, "World!") // true
	fmt.Println("Has suffix 'World!':", hasSuffix)

	index := strings.Index(s, "World") // 7
	fmt.Println("Index of 'World':", index)

	join := strings.Join([]string{"Hello", "World"}, ", ") // "Hello, World"
	fmt.Println("Joined string:", join)

	repeat := strings.Repeat("Go", 3) // "GoGoGo"
	fmt.Println("Repeated string:", repeat)

	replace := strings.Replace(s, "World", "Golang", 1) // "Hello, Golang!"
	fmt.Println("Replaced string:", replace)

	split := strings.Split(s, ", ") // ["Hello", "World!"]
	fmt.Println("Split string:", split)

	toLower := strings.ToLower(s) // "hello, world!"
	fmt.Println("Lowercase string:", toLower)

	toUpper := strings.ToUpper(s) // "HELLO, WORLD!"
	fmt.Println("Uppercase string:", toUpper)

	trimSpace := strings.TrimSpace("  Hello, World!  ") // "Hello, World!"
	fmt.Println("Trimmed string:", trimSpace)

	// 8. Функции пакета `strconv`
	intStr := "123"
	intValue, _ := strconv.Atoi(intStr) // 123
	fmt.Println("Converted int:", intValue)

	floatStr := "123.45"
	floatValue, _ := strconv.ParseFloat(floatStr, 64) // 123.45
	fmt.Println("Converted float:", floatValue)

	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr) // true
	fmt.Println("Converted bool:", boolValue)

	intToString := strconv.Itoa(456) // "456"
	fmt.Println("Int to string:", intToString)

	floatToString := strconv.FormatFloat(789.12, 'f', -1, 64) // "789.12"
	fmt.Println("Float to string:", floatToString)

	boolToString := strconv.FormatBool(false) // "false"
	fmt.Println("Bool to string:", boolToString)

	// 9. Функции пакета `unicode/utf8`
	runeCount := utf8.RuneCountInString(s) // Количество рун в строке
	fmt.Println("Rune count:", runeCount)

	rune, size := utf8.DecodeRuneInString(s) // Первая руна и её размер в байтах
	fmt.Println("First rune:", rune)
	fmt.Println("Rune size:", size)

}

package main

import (
	"fmt"
	"strings"
)

const (
	str    = "something"
	substr = "abc"
)

func main() {
	// String
	// 1. len string
	sentence := "Hello"
	lenSentence := len(sentence)
	fmt.Println(lenSentence)
	// 2. compare string
	str1 := "abc"
	str2 := "abd"
	fmt.Println(str1 == str2)
	// 3.contain
	res := strings.Contains(str, substr)
	fmt.Println(res)
	// 4. get value
	value := "cat;dog"
	substring := value[4:len(value)]
	fmt.Println(substring)
	// 5. replace
	s := "katakatakatak"
	// t := strings.Replace(s, "a", "o", -1) akan mengganti semua huruf a jika dibawah 0
	t := strings.Replace(s, "a", "o", 3) // mengganti 3 a pertama
	fmt.Println(t)
	// 6. insert
	p := "green"
	index := 2
	q := p[:index] + "Hi" + p[index:]
	fmt.Println(p)
	fmt.Println(q)

}

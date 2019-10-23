package main

import "fmt"

func main() {
	fmt.Println(comma2("januszNowak"))
}

//pierwsza iteracja
//input to "januszNowak", len parametru jest wieksze od 3 wiec wywolujemy comma(januszNo) ,wak zostaje odlozone na stos
//stos: ,wak
//druga iteracja
//input to januszNo, len jest wieksze od 3 wiec wywolujemy comma(janus) ,zNo zostaje odlozone na stos
//stos: 2),zNo 1),wak
//trzecia iteracja
//input to janus, len jest wieksze od 3 wiec wywolujemy comma(ja) ,nus zostaje odlozone na stos
//stos: 3),nus 2),zNo 1),wak
//czwarta iteracja
//input to ja, len jest mniejsze od 3 wiec zwracamy ja + rzeczy ze stosu
func comma(s string) string {
	l := len(s)
	if l <= 3 {
		return s
	}
	return comma(s[:l-3]) + "," + s[l-3:]
}

func comma2(s string) string {
	l := len(s)

	if l <= 3 {
		return s
	}
	return s[:3] + "," + comma2(s[3:])
}

func factorial(i int) int{
	if i == 1 {
		return 1
	}
	return i * factorial(i - 1)
}

func sum(i int) int {
	if i == 1 {
		return 1
	}
	return i + sum(i-1)
}

func reverse(s string) string {
	l := len(s)
	if l == 1 {
		return s
	}
	return reverse(s[1:]) + s[0:1]
}

func fib(i int) int {
	if i <= 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

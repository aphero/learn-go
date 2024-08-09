package main

import "fmt"

const (
	spanish         = "Spanish"
	french          = "French"
	german          = "German"
	enPrefix string = "Hello, "
	spPrefix string = "Hola, "
	frPrefix string = "Bonjour, "
	dePrefix string = "Hallo, "
)

func greetingPrefix(l string) (p string) {
	switch l {
	case spanish:
		p = spPrefix
	case french:
		p = frPrefix
	case german:
		p = dePrefix
	default:
		p = enPrefix
	}
	return
}

func Hello(n, l string) string {
	if n == "" {
		n = "world"
	}

	return greetingPrefix(l) + n
}

func main() {
	fmt.Println(Hello("world", ""))
}

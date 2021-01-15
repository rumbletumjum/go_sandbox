package main

import "fmt"
import x "awesomeProject/xor"

func main() {
	fmt.Println("hello")
	msg := "Hello Sofia"
	key := "fizzbutt"
	output := x.EncrypterDecrypter(msg, key)
	cleartext := x.EncrypterDecrypter(output, key)
	fmt.Printf("% x\n", output)
	fmt.Printf("%s\n", cleartext)
	fmt.Printf("%b\n", key)
	printMsg(cleartext, key)
	printBytes()
	fmt.Println(stringToBin(key))
}

func printBytes() {
	str := "abc"

	for i := range str {
		x := str[i]
		fmt.Printf("%b\n", x)
	}
}

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return binString
}

func printMsg(msg, key string) {
	fmt.Printf("% x\n", x.EncrypterDecrypter(msg, key))
}
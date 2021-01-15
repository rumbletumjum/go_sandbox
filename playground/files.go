package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Users struct {
	People []Person `json:"people"`
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Person) Write(w io.Writer) {
	b, _ := json.Marshal(*p)
	w.Write(b)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	mydata := []byte("the super important data")

	err := ioutil.WriteFile("file.dat", mydata, 0777)
	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("file.dat")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))

	// Open a file
	f, err := os.Create("dat1")
	check(err)
	defer f.Close()

	data2 := []byte("data for new file")
	n2, err := f.Write(data2)
	check(err)
	fmt.Printf("Wrote %d bytes\n", n2)

	f.Sync()

	me := Person{Id: 0, Name: "Ron", Age: 36}
	f2, err := os.Create("person.json")
	check(err)
	defer f2.Close()

	me.Write(f2)
}

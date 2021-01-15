package main

import (
    "bufio"
	"html/template"
    "io"
    "io/ioutil"
    "os"
)

type post struct {
    Title string
    Text string
}

type xrdb struct {
    Foreground string
    Background string
    Cursor string
    Colors []string
}
 
func (x *xrdb) initFromFile(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	x.Foreground = scanner.Text()
	scanner.Scan()
	x.Background = scanner.Text()
	scanner.Scan()
	x.Cursor = scanner.Text()

	for scanner.Scan() {
		x.Colors = append(x.Colors, scanner.Text())
	}
	return scanner.Err()
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    dat, _ := ioutil.ReadFile("xrdb.tmpl")
    t := template.Must(template.New("test").Parse(string(dat)))
    //p := post{"some page", "some text"}

    rdr := os.Stdin
    var x xrdb
    err := x.initFromFile(rdr)
    check(err)

    err = t.Execute(os.Stdout, x)
    check(err)
}

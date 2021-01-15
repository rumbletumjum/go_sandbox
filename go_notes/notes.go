package notes

import (
    "bufio"
    "fmt"
    "os"
)

func PublicFunc() {
    lines := readFile("/home/rkb/list")
    show(lines)
}

func readFile(filename string) []string {
    file, err := os.Open(filename)
    if err != nil {
        return []string{}
    }
    defer file.Close()

    scanner := bufio.NewScanner(bufio.NewReader(file))

    var tmp []string
    for scanner.Scan() {
        tmp = append(tmp, scanner.Text())
    }

    return tmp
}

func show(lines []string) {
    for _, line := range lines {
        fmt.Printf("%s\n", line)
    }
}

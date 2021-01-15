package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type response1 struct {
    Page int `json:"page"`
    Fruits []string `json:"fruits"`
}

type task struct {
    ID int `json:"_id"`
    Timestamp int `json:"timestamp"`
    Description string `json:"description"`
    Complete bool `json:"isComplete"`
    Boards []string `json:"boards"`
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    slcD := []string{"apple", "peach", "pear", "grapefruit"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    res1D := &response1{
        Page: 1,
        Fruits: slcD}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    str := `{"page":2,"fruits":["lingonberry","smultron"]}`
    res := response1{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    dat, err := ioutil.ReadFile("./tasks.json")
    check(err)
    tasks := []task{}
    json.Unmarshal(dat, &tasks)
    fmt.Println(tasks)
    fmt.Println(tasks[0].Complete)
}

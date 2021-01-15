package main

import (
    "net/http"
)

type CommandType int

const (
    GetCommand = iota
    SetCommand
    IncCommand
)

type Command struct {
    ct CommandType
    name string
    val int
    replyChan chan int
}

type Server struct {
    cmds chan<- Command
}

func startManagerGoroutine(initvals map[string]int) chan<- Command {
    counters := make(map[string]int)
    for k, v := range initvals {
        counters[k] = v
    }

    cmds := make(chan Command)

    go func() {
        for cmd := range cmds {
            switch cmd.ct {
            case GetCommand:
                if val, ok := counters[cmd.name]; ok {
                    cmd.replyChan <- val
                } else {
                    cmd.replyChan <- -1
                }
            case SetCommand:
                counters[cmd.name] = cmd.val
                cmd.replayChan <- val
            case IncCommand:
            default:
            }
        }
    }()

}

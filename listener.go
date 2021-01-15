package main

import (
    "io"
    "log"
    "net"
    "os"
)

func main() {
    listener, err := net.Listen("tcp", "0.0.0.0:56263")
    if err != nil {
            log.Fatal(err)
    }
    defer func() { _ = listener.Close() }()

    log.Printf("bound to %q", listener.Addr())

    for {
            conn, err := listener.Accept()
            if err != nil {
                    log.Fatal(err)
            }

            go func(c net.Conn) {
                defer c.Close()
                log.Printf("Connection from %s", c.RemoteAddr().String())
                file, err := os.Open("dummy.dat")
                if err != nil {
                        log.Fatal(err)
                }
                buf := make([]byte, 65536)

                written := 0
                for {
                    n, err := file.Read(buf)
                    if err == io.EOF {
                        break
                    }
                    c.Write(buf)
                    written += n
                    log.Printf("Bytes Written: %d", written)
                }

            }(conn)
    }
}


package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	buf := make([]byte, 512)
	for {
		log.Printf("Serving %s\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Fatalln(fmt.Errorf("error: %v", err))
			log.Fatalln("Unexpected error")
			break
		}
		log.Printf("Received %d bytes", n)

		log.Println("Writing data")
		if _, err := conn.Write(buf[0:]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatalln("Unable to bind port")
	}
	done := make(chan struct{})
	go func() {
		defer func() { done <- struct{}{} }()
		for {
			log.Printf("listening on %s\n", listener.Addr().String())
			conn, err := listener.Accept()

			if err != nil {
				log.Printf("this is the bogus-ish error: %v", err)
				return
			}

			go func(c net.Conn) {
				defer func() {
					c.Close()
					done <- struct{}{}
				}()
				buf := make([]byte, 1024)
				for {
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							log.Println(err)
						}
						return
					}
					log.Printf("received %d bytes: %q", n, buf[:n])
				}
			}(conn)
		}
	}()


	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		log.Println(err)
	}


	conn.Close()
	<-done
	listener.Close()
	<-done
}

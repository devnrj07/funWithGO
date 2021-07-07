//Package is about request and response handling at tcp level. An interesting example..package main

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue // because we still want to keep accepting connection
		}

		go handle(conn) // and pass those connections to go routines
	}
}

func handler(conn net.Conn) {
	defer conn.Close() // close connection after requesting handling is done

	//read the data
	request(conn)

	//write the data
	response(conn)
}

func request(conn net.Conn) {

	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		li := scanner.Text()
		fmt.Println()

		if i == 0 {
			method := strings.Fields(li)[0] //method
			uri := strings.Fields(li)[1]    // uri
			fmt.Println("**Method**", method)
			fmt.Println("**URI**", uri)
		}
		if li == "" {

			//headers are done. breakline followed by body
			break

		}
		i++
		// watch closely how we're iterating rows and columns
	}
}

func response(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Damn! That's hell lot of work</strong></body></html>`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length : %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type : text/html\r\n")
	fmt.Fprintf(conn, "\r\n") // break line to seperate header and body
	fmt.Fprintf(conn, body)
}

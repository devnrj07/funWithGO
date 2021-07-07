package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, error := net.Listen("tcp", ":8080") //watch closely how we define a port its a string not an int
	if error != nil {                             //handle error then and there
		log.Fatalln(error)
	}
	defer listener.Close() // Important step: to close the connection with defer keyword

	for {
		conn, err := listener.Accept() //we accept & an open connection has been established
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn) // we launch a go routine and pass our connection to it.
	}
}

func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn) // we will scan whatever is written by our connection

	for scanner.Scan() { //infite for loop with bool method(Scan()), it will break out of loop once it encounters a breakline
		scannedData := scanner.Text()
		fmt.Println(scannedData)
		fmt.Fprintf(conn, "Scanned msg %s \n", scannedData)

	}

	defer conn.Close()

	//we never get here
	// we have an open stream connection

	fmt.Println("code got here!!")
}

package main

import "fmt"

func main() {
	conn, err := net.Dial("tcp","localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	bs ,err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fataln(err)
	}

	fmt.Println(string(bs))
}
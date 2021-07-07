package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln("Fatal error", err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	if err != nil {
		return []byte{}, fmt.Errorf("ddf %v", err) //because it returns an error type
		//	return []byte{}, fmt.Printf("ddf %v", err)  //This won't work as returns errorstring Type
		//	return []byte{}, errors.New("ddf %v", err)  // This also doesn't worrks as it returns errorstring type
		//	return []byte{}, errors.New(fmt.Sprintf("ddf %v", err))) // This is a good solution Errorf is basically erros.New() + fmt.Sprintf()
	}
	return bs , nil // Important to return nil
}

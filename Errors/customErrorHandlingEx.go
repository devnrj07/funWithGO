package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	/*myError := sqrtError{
		lat:  "50.2289 N",
		long: "99.4656 W",
		err:  errors.New(fmt.Sprintf("Wow what the heck!")),
	}  This is just one way to make it more short. Lets do this
	*/

	if f < 0 {
		// write your error code here

		additionalError := errors.New(fmt.Sprintf("Wow what the heck! %v", f))
		//orError := fmt.Errorf("Wow what the heck! %v",f)

		return 0.0, sqrtError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  additionalError,
		}
	}
	return 42, nil
}

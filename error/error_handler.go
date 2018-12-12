package errorhandler

import (
	"fmt"
)

//HandleErr is a higher order function to simply handle errors
func HandleErr(err error, reason ...string) {
	if err != nil {
		fmt.Printf("Error %s :%s\n", err, reason)
		panic(err)
	}
}

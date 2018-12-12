package errorhandler

import (
	"fmt"
)

func HandleErr(err error, reasin ...string) {
	if err != nil {
		fmt.Printf("Error %s :%s\n", err, err)
		panic(err)
	}
}

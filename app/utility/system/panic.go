package system

import (
	"fmt"
	"log"
)

func SecurePanic(err error) {
	if err != nil {
		panic(err)
	}
}

// FatalError
func FatalError(msg string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s:%v", msg, err))
	}

	log.Printf("%s Success\n", msg)
}

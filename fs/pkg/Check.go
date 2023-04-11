package pkg

import (
	"fmt"
)

func Chekck(er error) {
	if er != nil {
		fmt.Println("Error in the input file!!!")
		return
	}
}

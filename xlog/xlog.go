package xlog

import (
	"fmt"
)

func Print(format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...) + "\n"
	fmt.Printf(str)
}
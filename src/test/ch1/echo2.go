package main

import (
	"fmt"
	"os"
	// "strings"
	"strconv"
)

func main(){
	s, sep := "", ""
	for i, arg := range os.Args[1:]{
		s += strconv.Itoa(i) + ". " + sep + arg
		sep = " "
	}
	fmt.Println(s)
	// fmt.Println(os.Args[1:])
	// fmt.Println(strings.Join(os.Args[1:], " "))
}
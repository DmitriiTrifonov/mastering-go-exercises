package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var sum, count float64
	args := os.Args
	for i, arg := range args {
		// args[0] is a program name arg
		if i == 0 {
			continue
		}
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Println(err)
			continue
		}
		sum += num
		count++
	}
	_, err := fmt.Fprintln(os.Stdout, sum/count)
	if err != nil {
		log.Println(err)
	}
}

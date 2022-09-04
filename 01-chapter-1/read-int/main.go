package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	readArgs := false
	arr := make([]int, 0)
	args := os.Args
	if len(args) > 1 {
		readArgs = true
	}
	if readArgs {
		for i, arg := range args {
			if i == 0 {
				continue
			}
			if arg == "END" {
				break
			}
			num, err := strconv.ParseInt(arg, 10, 64)
			if err != nil {
				log.Println(err)
				continue
			}
			arr = append(arr, int(num))
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			str := scan.Text()
			if str == "END" {
				break
			}
			num, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				log.Println(err)
				continue
			}
			arr = append(arr, int(num))
		}
	}
	_, err := fmt.Fprintln(os.Stdout, arr)
	if err != nil {
		log.Println(err)
	}

}

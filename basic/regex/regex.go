package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is 740652552@163.com@abc.com
email2 is abc@ef.com
email3 is sdsadsa@sdadsa.comsadsadsds.com
`

func main() {
	re  := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
	for _, value := range match {
		for _, v := range  value {
			fmt.Print(v)
			fmt.Print("--------")
		}
		fmt.Println()
	}
}

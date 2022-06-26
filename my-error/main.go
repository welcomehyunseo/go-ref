package main

import (
	"fmt"
	"myerror/package1"
	"myerror/package1/package2"
	"myerror/package1/package2/package3"
)

func main() {
	if err := package1.MakeError(); err != nil {
		fmt.Println(err.Error())
	}
	if err := package2.MakeError(); err != nil {
		fmt.Println(err.Error())
	}
	if err := package3.MakeError(); err != nil {
		fmt.Println(err.Error())
	}
}

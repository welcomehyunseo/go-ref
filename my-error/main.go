package main

import (
	"myerror/package1"
	"myerror/package1/package2"
	"myerror/package1/package2/package3"
)

func main() {
	package1.MakeError()
	package2.MakeError()
	package3.MakeError()
}

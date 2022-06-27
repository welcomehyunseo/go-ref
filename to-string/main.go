package main

import "fmt"

type Box interface {
	Area() int
}

type CoolBox struct {
	width  int
	height int
}

func (b *CoolBox) Area() int {
	return b.width * b.height
}

func (b *CoolBox) String() string {
	return fmt.Sprintf("CoolBox --->>>>> (Width: %d, Height: %d)", b.width, b.height)
}

func main() {
	coolBox := &CoolBox{
		width:  10,
		height: 10,
	}
	fmt.Println(coolBox)

	var box Box = &CoolBox{
		width:  20,
		height: 30,
	}
	fmt.Printf("%s", box)
}

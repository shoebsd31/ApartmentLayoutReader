package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

type LayoutStructure struct {
	Corners int
	Height  int
	Width   int
}

func main() {

	content, err := ioutil.ReadFile("rooms.txt")
	//var layout LayoutStrcture

	if err != nil {
		log.Fatal(err)
	}
	layout, err1 := HandleContent(string(content))

	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(layout.Corners)
	fmt.Println(string(content))

}

func HandleContent(content string) (LayoutStructure, error) {

	var cornersCount = 0
	first := content[0:1]
	last := content[len(content)-1:]

	if first != "+" {
		return LayoutStructure{}, errors.New("apartment doesn't start with +")
	}
	if last != "+" {
		return LayoutStructure{}, errors.New("apartment doesn't end with +")
	}
	// for _, x := range content {

	// 	if x == '+' {
	// 		cornersCount++
	// 	}

	// }
	for i := 0; i < len(content); i++ {
		if string(content[i]) == "+" {
			cornersCount++
		}
	}
	return LayoutStructure{
		Corners: cornersCount,
		Width:   len(content),
	}, nil
}

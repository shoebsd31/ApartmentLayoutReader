package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type LayoutStructure struct {
	Corners      int
	Height       int
	MinimumWidth int
	MaximumWidth int
	RoomName     string
}

func main() {

	content, err := ioutil.ReadFile("rooms.txt")

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
	var widthCount = 0
	first := content[0:1]
	last := content[len(content)-1:]
	// var height = 0
	var wallsize = 0
	if first != "+" {
		return LayoutStructure{}, errors.New("apartment doesn't start with +")
	}
	if last != "+" {
		return LayoutStructure{}, errors.New("apartment doesn't end with +")
	}

	var newLines = strings.Split(content, "\n\t")
	var isRoomNameFound = false
	var roomStarts = false
	var roomname string = ""
	var cornerFound = false
	var widthArray []int
	for i := 0; i < len(content); i++ {

		if string(content[i]) == "+" {
			cornersCount++
			cornerFound = true
			//another end is found
			if cornersCount%2 == 0 {
				cornerFound = false
				widthArray = append(widthArray, widthCount)
				widthCount = 0
			}
		}

		// if string(content[i]) == "\n" {
		// 	height++
		// }
		if cornerFound && string(content[i]) == "-" {
			widthCount++
		}
		//wall found
		if string(content[i]) == "|" || string(content[i]) == "/" || string(content[i]) == "\\" {
			wallsize++
		}

		var sb strings.Builder

		if roomStarts && isRoomNameFound {
			sb.WriteString(roomname)
			if string(content[i]) != ")" {
				sb.WriteString(string(content[i]))
			}
			roomname = sb.String()
		}
		if string(content[i]) == "(" {
			isRoomNameFound = true
			roomStarts = true
		}
		if string(content[i]) == ")" {
			isRoomNameFound = true
			roomStarts = false
		}
	}
	min, max := findMinAndMax(widthArray)
	return LayoutStructure{
		Corners:      cornersCount,
		Height:       len(newLines),
		RoomName:     roomname,
		MinimumWidth: min,
		MaximumWidth: max,
	}, nil
}

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

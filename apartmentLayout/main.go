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
	ChairArray   []string
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
	var height = 0
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
	var chairArray []string
	var wallStart = false
	for i := 0; i < len(content); i++ {

		if string(content[i]) == "\t" {
			continue
		}

		if string(content[i]) == "\n" {
			height++
			continue
		}

		if string(content[i]) == "+" {
			cornersCount++
			cornerFound = true
			//another end is found
			if cornersCount%2 == 0 {
				cornerFound = false
				widthArray = append(widthArray, widthCount)
				widthCount = 0
			}
			continue
		}

		if cornerFound && string(content[i]) == "-" {
			widthCount++
			continue
		}
		//wall found
		if string(content[i]) == "|" || string(content[i]) == "/" || string(content[i]) == "\\" {
			wallsize++
			wallStart = !wallStart
			continue
		}

		var sb strings.Builder

		if roomStarts && isRoomNameFound {
			sb.WriteString(roomname)
			if string(content[i]) != ")" {
				sb.WriteString(string(content[i]))
			} else {
				roomStarts = !roomStarts
				continue
			}
			roomname = sb.String()
		}
		if string(content[i]) == "(" {
			isRoomNameFound = true
			roomStarts = !roomStarts
			continue
		}

		if !roomStarts && string(content[i]) != " " {
			chairArray = append(chairArray, string(content[i]))
		}

	}
	min, max := findMinAndMax(widthArray)
	return LayoutStructure{
		Corners:      cornersCount,
		Height:       len(newLines),
		RoomName:     roomname,
		MinimumWidth: min,
		MaximumWidth: max,
		ChairArray:   chairArray,
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

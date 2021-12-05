package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func Test_CheckCorners(t *testing.T) {

	result, err := HandleContent("+--------+-------+")
	if err != nil {
		t.Errorf("Handle content failed with error %v", err)
	}
	t.Logf("Total Corners SUCCESS, expected %d, got %d", 3, result.Corners)
	t.Logf("Total Width SUCCESS, expected %d, got %d", 18, result.MaximumWidth)

}

func Test_ValidateCorners(t *testing.T) {

	_, err := HandleContent("--------+-------+")
	if err != nil {
		var errMessage = "apartment doesn't start with +"
		t.Logf("Expected '%v', GOT  '%v'", err, errMessage)
	}

	_, err2 := HandleContent("+--------+-------")
	if err2 != nil {
		var errMessage = "apartment doesn't end with +"
		t.Logf("Expected '%v', GOT  '%v'", err, errMessage)
	}

}

func Test_ValidateWidthCharacter(t *testing.T) {

	_, err := HandleContent("--------+-------+")
	if err != nil {
		var errMessage = "apartment doesn't start with +"
		t.Logf("Expected '%v', GOT  '%v'", err, errMessage)
	}

	_, err2 := HandleContent("+--------+-------")
	if err2 != nil {
		var errMessage = "apartment doesn't end with +"
		t.Logf("Expected '%v', GOT  '%v'", err, errMessage)
	}
}

func Test_ReadRoom(t *testing.T) {
	layout, err := HandleContent(`+-----------+
	|           |
	| (closet)  |
	|         P |
	|         P |
	|         P |
	|           |
	+-----------+`)

	var expectedCorners = 4
	var expectedWidth = 11
	var expectedHeight = 8
	var expectedRoomName = "closet"
	var expectedChairArray []string
	expectedChairArray = append(expectedChairArray, "P")
	expectedChairArray = append(expectedChairArray, "P")
	expectedChairArray = append(expectedChairArray, "P")
	if err != nil {
		t.Errorf("FAILED with error: '%v'", err)
	}

	if layout.Corners == expectedCorners {
		t.Logf("Expected Corners '%v', GOT  '%v'", expectedCorners, layout.Corners)
	} else {
		t.Errorf("Expected Corners '%v', GOT  '%v'", expectedCorners, layout.Corners)
	}

	if layout.MaximumWidth == expectedWidth {
		t.Logf("Expected Width '%v', GOT  '%v'", expectedWidth, layout.MaximumWidth)
	} else {
		t.Errorf("Expected Width '%v', GOT  '%v'", expectedWidth, layout.MaximumWidth)
	}

	if layout.Height == expectedHeight {
		t.Logf("Expected Height '%v', GOT  '%v'", expectedHeight, layout.Height)
	} else {
		t.Errorf("Expected Height '%v', GOT  '%v'", expectedHeight, layout.Height)
	}

	if layout.RoomName == expectedRoomName {
		t.Logf("Expected Room name '%v', GOT  '%v'", expectedRoomName, layout.RoomName)
	} else {
		t.Errorf("Expected Room name '%v', GOT  '%v'", expectedRoomName, layout.RoomName)
	}

	if cmp.Equal(layout.ChairArray, expectedChairArray) {

		for index, itr := range layout.ChairArray {
			t.Logf("Expected ChairArray '%v' at '%d', GOT  '%v'", itr, index, expectedChairArray[index])
		}

	} else {
		t.Errorf("Expected ChairArray '%v', GOT  '%v'", expectedRoomName, layout.RoomName)
	}
}

package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func Test_CheckCorners(t *testing.T) {

	result, err := HandleContent("+--------+-------+")
	if err != nil {
		t.Errorf("Handle content failed with error %v", err)
	}
	t.Logf("Total Corners SUCCESS, expected %d, got %d", 3, result.Corners)
	t.Logf("Total Width SUCCESS, expected %d, got %d", 18, result.Width)

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

	_, err := HandleContent(`+-----------+
	|           |
	| (closet)  |
	|         P |
	|         P |
	|         P |
	|           |
	+-----------+`)

	if err != nil {
		t.Errorf("FAILED with error: '%v'", err)
	}

}

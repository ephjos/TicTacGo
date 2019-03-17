package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/josephthomashines/TicTacGo/board"
)

func TestBoardBuild(t *testing.T) {
	b := board.Init(3, 3)
	bp := &b

	actual := bp.Get()

	expected := make([][]string, 3)

	for i := range expected {
		expected[i] = make([]string, 3)
	}

	for i := range expected {
		for j := range expected[i] {
			expected[i][j] = board.PLACEHOLDER
		}
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nExpected \n%v\n to actually be \n%v", expected, actual)
	}
}

func TestBoardPrint(t *testing.T) {
	b := board.Init(3, 3)
	bp := &b

	actual := fmt.Sprint(bp)

	expected := "_ _ _ \n_ _ _ \n_ _ _ \n"

	if actual != expected {
		t.Errorf("\nExpected\n%v\n to actually be \n%v", expected, actual)
	}
}

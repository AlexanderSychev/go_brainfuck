package syntax

import (
	"reflect"
	"testing"
)

func TestCanonize(t *testing.T) {
	t.Log("\"canonize\" function test")

	src := "++++++++++\n\t[>+++++++>++++++++++>+++>+<<<<-\n]\r\n>++.>+.+++++++..+++.>++.  <<+++++++++++++++.>.+++.------.--------.>+.>."
	expected := "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."

	result := canonize(src)
	if result != expected {
		t.Errorf("Expected \"%s\" got \"%s\"", expected, result)
	}
}

func TestExplode(t *testing.T) {
	t.Log("\"explode\" function test")

	src := ">+++++++>++++++++++>+++>+<<<<-"
	expected := []string{
		">",
		"+",
		"+",
		"+",
		"+",
		"+",
		"+",
		"+",
		">",
		"+",
		"+",
		"+",
		"+",
		"+",
		"+",
		"+",
		"+",
		"+",
		"+",
		">",
		"+",
		"+",
		"+",
		">",
		"+",
		"<",
		"<",
		"<",
		"<",
		"-",
	}

	result := explode(src)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected \"%v\" got \"%v\"", expected, result)
	}
}

func TestCheckCycles(t *testing.T) {
	t.Log("\"checkCycles\" function tests")

	t.Log("Code with not opened cycle")
	src01 := "]++[.]"
	expected01 := "Syntax error: unexpected token \"]\""
	err01 := checkCycles(src01)
	if err01 == nil || err01.Error() != expected01 {
		t.Errorf("Expected error with message \"%s\"", expected01)
	}

	t.Log("Code with not closed cycle")
	src02 := "[--++]++[."
	expected02 := "Syntax error: unexpected token \"[\""
	err02 := checkCycles(src02)
	if err02 == nil || err02.Error() != expected02 {
		t.Errorf("Expected error with message \"%s\"", expected02)
	}
}

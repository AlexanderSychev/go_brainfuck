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

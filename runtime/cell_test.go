package runtime

import "testing"

func TestCell_GetValue(t *testing.T) {
	t.Log("\"GetValue()\" method test")

	var expected byte = 5
	cell := Cell{value: expected}
	if cell.GetValue() != expected {
		t.Errorf("Expected %d got %d", expected, cell.GetValue())
	}
}

func TestCell_SetValue(t *testing.T) {
	t.Log("\"SetValue()\" method test")

	var expected byte = 6
	cell := Cell{value: 5}

	cell.SetValue(expected)
	if cell.GetValue() != expected {
		t.Errorf("Expected %d got %d", expected, cell.GetValue())
	}
}

func TestCell_Next(t *testing.T) {
	t.Log("\"Next()\" method test")

	t.Log("If no next cell defined then method creates it")
	cell01 := NewZeroCell()
	next01 := cell01.Next()
	if next01 == nil {
		t.Error("Expected \"Cell\" object got \"nil\"")
	} else if next01 != cell01.next {
		t.Error("\"next\" property must equal result of \"Next()\" method")
	}

	t.Log("If next cell defined then method returns it")
	cell02 := NewZeroCell()
	next02 := NewZeroCell()
	cell02.next = next02
	next02.previous = cell02
	if cell02.Next() == nil {
		t.Error("Expected \"Cell\" object got \"nil\"")
	} else if cell02.Next() != next02 {
		t.Error("\"next02\" must equal result of \"Next()\" method")
	}
}

func TestCell_Previous(t *testing.T) {
	t.Log("\"Previous()\" method test")

	t.Log("If no previous cell defined then method creates it")
	cell01 := NewZeroCell()
	previous01 := cell01.Previous()
	if previous01 == nil {
		t.Error("Expected \"Cell\" object got \"nil\"")
	} else if previous01 != cell01.previous {
		t.Error("\"previous\" property must equal result of \"Next()\" method")
	}

	t.Log("If next cell defined then method returns it")
	cell02 := NewZeroCell()
	previous02 := NewZeroCell()
	cell02.previous = previous02
	previous02.next = cell02
	if cell02.Previous() == nil {
		t.Error("Expected \"Cell\" object got \"nil\"")
	} else if cell02.Previous() != previous02 {
		t.Error("\"next\" property must equal result of \"Next()\" method")
	}
}

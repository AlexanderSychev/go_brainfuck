package runtime

// Memory contains pointer
type Memory struct {
	currentCell *Cell
}

// SetValue updates value of current cell
func (m *Memory) SetValue(value byte) {
	m.currentCell.SetValue(value)
}

// GetValue returns value of current cell
func (m Memory) GetValue() byte {
	return m.currentCell.GetValue()
}

// ToNext set next cell as current
func (m *Memory) ToNext() {
	m.currentCell = m.currentCell.Next()
}

// ToPrevious set previous cell as current
func (m *Memory) ToPrevious() {
	m.currentCell = m.currentCell.Previous()
}

func NewMemory() *Memory {
	return &Memory{
		currentCell: NewZeroCell(),
	}
}

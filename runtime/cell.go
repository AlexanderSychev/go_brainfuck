package runtime

// Cell is elementary unit of runtime which contains single value and pointers to neighbour cells
type Cell struct {
	// value of cell (byte representation of ASCII character)
	value byte
	// previous cell pointer
	previous *Cell
	// next cell pointer
	next *Cell
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Methods
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// GetValue return current value of cell
func (c Cell) GetValue() byte {
	return c.value
}

// SetValue set new value of cell
func (c *Cell) SetValue(value byte) {
	c.value = value
}

// Next return pointer to next cell. If next cell is not exists then method creates it.
func (c *Cell) Next() *Cell {
	if c.next == nil {
		c.next = &Cell{
			value:    0,
			previous: c,
			next:     nil,
		}
	}

	return c.next
}

// Previous return pointer to previous cell. If previous cell is not exists then method creates it.
func (c *Cell) Previous() *Cell {
	if c.previous == nil {
		c.previous = &Cell{
			value:    0,
			previous: nil,
			next:     c,
		}
	}

	return c.previous
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Constructors
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// NewZeroCell creates new Cell instance with zero value and no links
func NewZeroCell() *Cell {
	return &Cell{
		value:    0,
		previous: nil,
		next:     nil,
	}
}

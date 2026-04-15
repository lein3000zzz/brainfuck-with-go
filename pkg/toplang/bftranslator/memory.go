package bftranslator

// memory manages the brainfuck tape cell allocation
// each variable is assigned a dedicated cell. Temporary cells are used
// for intermediate arithmetic results and are recycled via a free-list
type memory struct {
	nextCell  int // bump allocator for fresh cells
	maxCells  int
	variables map[string]int // variable name - cell index
	freeTemps []int          // recycled temp cells
}

func newMemory(maxCells int) *memory {
	return &memory{
		nextCell:  0,
		maxCells:  maxCells,
		variables: make(map[string]int),
		freeTemps: make([]int, 0, maxCells),
	}
}

// declareVar assigns the next free cell to a variable. Returns the cell index
func (m *memory) declareVar(name string) (int, error) {
	if _, exists := m.variables[name]; exists {
		return 0, ErrRedeclaredVariable
	}

	cell, err := m.allocCell()
	if err != nil {
		return 0, err
	}

	m.variables[name] = cell
	return cell, nil
}

// lookupVar returns the cell index of a declared variable
func (m *memory) lookupVar(name string) (int, error) {
	cell, ok := m.variables[name]
	if !ok {
		return 0, ErrUndeclaredVariable
	}

	return cell, nil
}

// allocTemp returns a temporary cell (reuses freed ones first)
func (m *memory) allocTemp() (int, error) {
	if len(m.freeTemps) > 0 {
		cell := m.freeTemps[len(m.freeTemps)-1]
		m.freeTemps = m.freeTemps[:len(m.freeTemps)-1]

		return cell, nil
	}

	return m.allocCell()
}

// freeTemp returns a temporary cell to the free-list for reuse
func (m *memory) freeTemp(cell int) {
	m.freeTemps = append(m.freeTemps, cell)
}

// allocCell bumps the next-cell pointer.
func (m *memory) allocCell() (int, error) {
	if m.nextCell >= m.maxCells {
		return 0, ErrOutOfMemory
	}

	cell := m.nextCell
	m.nextCell++
	return cell, nil
}

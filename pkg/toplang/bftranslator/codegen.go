package bftranslator

import "strings"

// codegen emits low-level brainfuck instruction sequences
// It tracks the current tape head position so moveTo emits minimal moves
type codegen struct {
	out     []byte
	headPos int
}

func (c *codegen) moveTo(cell int) {
	diff := cell - c.headPos
	if diff > 0 {
		c.out = append(c.out, []byte(strings.Repeat(">", diff))...)
	} else if diff < 0 {
		c.out = append(c.out, []byte(strings.Repeat("<", -diff))...)
	}

	c.headPos = cell
}

func (c *codegen) emit(s string) {
	c.out = append(c.out, []byte(s)...)
}

// clearCell: cell = 0.
func (c *codegen) clearCell(cell int) {
	c.moveTo(cell)
	c.emit("[-]")
}

// setCell: cell = value (0-255).
func (c *codegen) setCell(cell int, value int) {
	c.clearCell(cell)
	if value > 0 {
		c.emit(strings.Repeat("+", value%256))
	}
}

// moveValue: dst += src; src = 0 (destructive move / add)
func (c *codegen) moveValue(src, dst int) {
	c.moveTo(src)
	c.emit("[")
	c.moveTo(dst)
	c.emit("+")
	c.moveTo(src)
	c.emit("-]")
}

// copyCell: dst = src (non-destructive). dst and tmp must be 0 before call.
// Pattern: src[dst+ tmp+ src-] tmp[src+ tmp-]
func (c *codegen) copyCell(src, dst, tmp int) {
	c.moveTo(src)
	c.emit("[")
	c.moveTo(dst)
	c.emit("+")
	c.moveTo(tmp)
	c.emit("+")
	c.moveTo(src)
	c.emit("-]")
	c.moveValue(tmp, src) // restore src
}

// addCells: result = a + b.
func (c *codegen) addCells(a, b, result int) {
	c.moveValue(a, result)
	c.moveValue(b, result)
}

// subCells: result = a - b (wrapping 8-bit).
func (c *codegen) subCells(a, b, result int) {
	c.moveValue(a, result)
	c.moveTo(b)
	c.emit("[")
	c.moveTo(result)
	c.emit("-")
	c.moveTo(b)
	c.emit("-]")
}

// multiplyCells: result = a * b. tmp must be 0.
// a[ b[result+ tmp+ b-] tmp[b+ tmp-] a- ]
func (c *codegen) multiplyCells(a, b, result, tmp int) {
	c.moveTo(a)
	c.emit("[")
	c.moveTo(b)
	c.emit("[")
	c.moveTo(result)
	c.emit("+")
	c.moveTo(tmp)
	c.emit("+")
	c.moveTo(b)
	c.emit("-]")
	c.moveValue(tmp, b)
	c.moveTo(a)
	c.emit("-]")
}

// normalizeBool: cell = (cell != 0). tmp must be 0.
// Pattern: cell[tmp + cell[-]] tmp[cell + tmp-]
func (c *codegen) normalizeBool(cell, tmp int) {
	c.moveTo(cell)
	c.emit("[")
	c.moveTo(tmp)
	c.emit("+")
	c.moveTo(cell)
	c.emit("[-]]")
	c.moveValue(tmp, cell)
}

// emitIfElse: if condCell != 0 - ifFn, else - elseFn.
// condCell is consumed. tmp must be 0.
func (c *codegen) emitIfElse(condCell, tmp int, ifFn func(), elseFn func()) {
	c.setCell(tmp, 1)
	c.moveTo(condCell)
	c.emit("[")

	if ifFn != nil {
		ifFn()
	}

	c.moveTo(tmp)
	c.emit("-")
	c.moveTo(condCell)
	c.emit("[-]]")
	c.moveTo(tmp)
	c.emit("[")

	if elseFn != nil {
		elseFn()
	}

	c.moveTo(tmp)
	c.emit("[-]]")
}

// emitLessThan: result = (a < b). a and b consumed
// Uses 3 scratch cells (t1, t2, t3) that must be 0.
//
// Algorithm - parallel decrement:
//
//	while a > 0:
//	  if b > 0: a--; b--
//	  else:     a = 0         // a ≥ b, force exit
//	result = bool(b)          // if b still > 0 after loop, a was smaller
func (c *codegen) emitLessThan(a, b, result, t1, t2, t3 int) {
	c.clearCell(result)

	c.moveTo(a)
	c.emit("[")

	c.clearCell(t1)
	c.clearCell(t2)
	c.copyCell(b, t1, t2)

	c.clearCell(t2)
	c.normalizeBool(t1, t2)

	c.setCell(t3, 1) // else-flag
	c.moveTo(t1)
	c.emit("[")

	c.moveTo(a)
	c.emit("-")
	c.moveTo(b)
	c.emit("-")
	c.moveTo(t3)
	c.emit("-")
	c.moveTo(t1)
	c.emit("[-]]")

	c.moveTo(t3)
	c.emit("[")
	c.moveTo(a)
	c.emit("[-]")
	c.moveTo(t3)
	c.emit("[-]]")

	c.moveTo(a)
	c.emit("]")

	c.clearCell(t1)
	c.normalizeBool(b, t1)
	c.moveValue(b, result)
}

// emitDivMod: quotient = a / b, remainder = a % b.
// a and b consumed. All output - scratch cells must be 0.
//
// Algorithm (repeated subtraction, does NOT modify remainder until confirmed):
//
//	remainder = a; quotient = 0;
//	outer_flag = 1
//	while outer_flag:
//	  outer_flag = 0
//	  rcopy = copy(remainder)   // working copy for trial subtraction
//	  counter = copy(b)
//	  success = 1
//	  while counter:
//	    if rcopy > 0: rcopy--; counter--
//	    else:         success = 0; counter = 0
//	  if success:
//	    remainder -= b          // apply the subtraction for real
//	    quotient++; outer_flag = 1
//	  (else: do nothing, remainder untouched, stop)
func (c *codegen) emitDivMod(a, b, quotient, remainder, rcopy, counter, success, t1, t2, t3 int) {
	c.clearCell(remainder)
	c.moveValue(a, remainder)
	c.clearCell(quotient)

	c.setCell(t1, 1)
	c.moveTo(t1)
	c.emit("[")
	c.clearCell(t1)

	c.clearCell(rcopy)
	c.clearCell(t2)
	c.copyCell(remainder, rcopy, t2)

	c.clearCell(counter)
	c.clearCell(t2)
	c.copyCell(b, counter, t2)

	c.setCell(success, 1)

	c.moveTo(counter)
	c.emit("[")

	c.clearCell(t2)
	c.clearCell(t3)
	c.copyCell(rcopy, t2, t3)
	c.clearCell(t3)
	c.normalizeBool(t2, t3)

	c.setCell(t3, 1)
	c.moveTo(t2)
	c.emit("[")

	c.moveTo(rcopy)
	c.emit("-")
	c.moveTo(counter)
	c.emit("-")
	c.moveTo(t3)
	c.emit("-")
	c.moveTo(t2)
	c.emit("[-]]")

	c.moveTo(t3)
	c.emit("[")
	c.clearCell(success)
	c.clearCell(counter)
	c.moveTo(t3)
	c.emit("[-]]")

	c.moveTo(counter)
	c.emit("]")

	c.clearCell(t2)
	c.emitIfElse(success, t2,
		func() {
			c.clearCell(t3)
			c.clearCell(t2) // scratch for copyCell
			c.copyCell(b, t3, t2)

			c.moveTo(t3)
			c.emit("[")
			c.moveTo(remainder)
			c.emit("-")
			c.moveTo(t3)
			c.emit("-]")

			c.moveTo(quotient)
			c.emit("+")
			c.setCell(t1, 1) // continue outer loop
		},
		nil, // else: do nothing
	)

	// end outer while
	c.moveTo(t1)
	c.emit("]")

	// clean up b (consumed per contract)
	c.clearCell(b)
	// clean up rcopy
	c.clearCell(rcopy)
}

// String returns the accumulated brainfuck program.
func (c *codegen) String() string {
	return string(c.out)
}

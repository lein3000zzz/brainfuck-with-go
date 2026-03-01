package astbuilder

// Посматриваю сюда: https://gist.github.com/roachhd/dce54bec8ba55fb17d3a

const (
	// IncreasePointerSign increases memory pointer, or moves the pointer to the right 1 block
	IncreasePointerSign = '>'

	// DecreasePointerSign decreases memory pointer, or moves the pointer to the left 1 block.
	DecreasePointerSign = '<'

	// IncreaseValueByPointerSign increases value stored at the block pointed to by the memory pointer
	IncreaseValueByPointerSign = '+'

	// DecreaseValueByPointerSign decreases value stored at the block pointed to by the memory pointer
	DecreaseValueByPointerSign = '-'

	// WhileLoopStartSign like c while(cur_block_value != 0) loop.
	WhileLoopStartSign = '['

	// WhileLoopEndSign if block currently pointed to's value is not zero, jump back to [
	WhileLoopEndSign = ']'

	// InputOneCharSign like c getchar(). input 1 character.
	InputOneCharSign = ','

	// PrintOneCharSign like c putchar(). print 1 character to the console
	PrintOneCharSign = '.'
)

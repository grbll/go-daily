package sudokusolver

type indLink struct {
	val  int
	next *indLink
	prev *indLink
}

type checker struct {
	rowMap []map[byte]struct{} // will store whether a byte can be placed in the corresponding row
	colMap []map[byte]struct{} // ...col
	squMap []map[byte]struct{} // ... squ
	curInd *indLink            //stores the order of the insert positions, where curInd.next will be the next index that needs to be filled by us
}

func NewChecker(board [][]byte) *checker {

	//used temp vars
	var row, col int
	var first, last, head *indLink

	//prepare first and head
	first, last = &indLink{-1, nil, nil}, &indLink{82, nil, nil}
	first.next, last.prev = last, first
	head = first

	//initialize new checker
	var ch checker = checker{rowMap: make([]map[byte]struct{}, 9, 9), colMap: make([]map[byte]struct{}, 9, 9), squMap: make([]map[byte]struct{}, 9, 9), curInd: first}

	//initialize all maps
	for i := 0; i < 9; i++ {
		ch.rowMap[i], ch.colMap[i], ch.squMap[i] = make(map[byte]struct{}), make(map[byte]struct{}), make(map[byte]struct{})
	}

	//build knowledge about bytes
	for ind := 0; ind < 81; ind++ {
		row, col = ind/9, ind%9
		if board[row][col] != '.' { //set original byte if non-empty
			ch.set(row, col, board[row][col])
		} else { //link possible insert position
			head.next.prev = &indLink{val: ind, next: head.next, prev: head}
			head.next = head.next.prev
			head = head.next
		}
	}
	ch.curInd = ch.curInd.next
	return &ch
}

// checks whether the placement of `b` at (`row`,`col`) conflicts with previous placements
func (ch *checker) check(row, col int, b byte) bool {
	if _, ok := ch.rowMap[row][b]; ok {
		return false
	}
	if _, ok := ch.colMap[col][b]; ok {
		return false
	}
	if _, ok := ch.squMap[(row/3)*3+(col/3)][b]; ok {
		return false
	}
	return true
}

// sets `b` at (`row`,`col`) if possible, return succsess value
func (ch *checker) set(row, col int, b byte) {
	ch.rowMap[row][b] = struct{}{}
	ch.colMap[col][b] = struct{}{}
	ch.squMap[(row/3)*3+(col/3)][b] = struct{}{}
}

// delets knowledge about `b` at (`row`,`col`)
func (ch *checker) del(row, col int, b byte) {
	delete(ch.rowMap[row], b)
	delete(ch.colMap[col], b)
	delete(ch.squMap[(row/3)*3+(col/3)], b)
}

// gets the next candidate which should be tried based on last attempt return '.' if no possible next
var nextCanditateMap map[byte]byte = map[byte]byte{'.': '1', '1': '2', '2': '3', '3': '4', '4': '5', '5': '6', '6': '7', '7': '8', '8': '9', '9': '.'}

func solveSudoku(board [][]byte) {
	//declare used vars
	var ch *checker  //contains conflict info about board, as well as insert position information
	var row, col int //temp stores to convert index to raster position
	var cur byte     //temp stores attempt of placement

	//initialize checker
	ch = NewChecker(board)

	for -1 < ch.curInd.val && ch.curInd.val < 82 { //while the current insert index is within the board ...
		row, col = ch.curInd.val/9, ch.curInd.val%9

		cur = nextCanditateMap[board[row][col]] //initialize next attempt i.e. if current '.' next '1', if current '2' -> next '3', if current '9' next '.'
		if cur != '1' {                         //last value was different from '.'
			ch.del(row, col, board[row][col]) //delete information about last value
		}

		for cur != '.' && !ch.check(row, col, cur) { //find next valid insert, or end at cur = '.'
			cur = nextCanditateMap[cur]
		}

		board[row][col] = cur //set board at (`row`,`col`) to new value

		if board[row][col] == '.' { //if new value is '.' (i.e. nothing worked)...
			ch.curInd = ch.curInd.prev //go back.
		} else { //Otherwise...
			ch.set(row, col, board[row][col]) //set the confilct informaton about the new value and ...
			ch.curInd = ch.curInd.next        //go to the next insert index.
		}
	}
}

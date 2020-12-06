package models

type Partitioner func(string) (int, int)

type BoardingPass struct {
	code string
	row  int
	col  int
	id   int
}

func NewBoardingPass(code string) *BoardingPass {
	return &BoardingPass{
		code: code,
	}
}

func (bp *BoardingPass) Partition(p Partitioner) {
	bp.row, bp.col = p(bp.code)
	bp.id = 8*bp.row + bp.col
}

func (bp BoardingPass) Seat() (int, int) {
	return bp.row, bp.col
}

func (bp BoardingPass) SeatID() int {
	return bp.id
}

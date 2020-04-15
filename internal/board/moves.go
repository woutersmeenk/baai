package board

type initiateDraw struct{}

type draw struct {
	tile tile
}

type reject struct{}

type place struct {
	tile tile
	pos  position
	farm bool
}

type claim struct {
	pos position
}

type moveExecuter interface {
	execute(b *board)
}

func (m initiateDraw) execute(b *board) { b.initiateDraw() }
func (m draw) execute(b *board)         { b.draw(m.tile) }
func (m reject) execute(b *board)       { b.reject() }
func (m place) execute(b *board)        { b.place(m.pos, m.tile, m.farm) }
func (m claim) execute(b *board)        { b.claim(m.pos) }

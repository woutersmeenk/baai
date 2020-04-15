package board

import "fmt"

func (b *board) nextPlayer() {
	switch b.currentPlayer {
	case one:
		b.currentPlayer = two
	case two:
		b.currentPlayer = one
	default:
		panic(fmt.Sprintf("Invalid player: %v", b.currentPlayer))
	}

	b.turnPhase = start
}

func (b *board) assertTurnPhase(phase turnPhase) {
	if b.turnPhase != phase {
		panic(fmt.Sprintf("Invalid turn phase: %v expected %v", b.turnPhase, phase))
	}
}

func (b *board) assertRoundPhase(phase roundPhase) {
	if b.roundPhase != phase {
		panic(fmt.Sprintf("Invalid round phase: %v expected %v", b.turnPhase, phase))
	}
}

func (b *board) initiateDraw() {
	b.assertRoundPhase(tiling)
	b.assertTurnPhase(start)

	b.turnPhase = drawingTile
}

func (b *board) draw(tile tile) {
	b.assertRoundPhase(tiling)
	b.assertTurnPhase(drawingTile)
	b.closedTiles[tile]--

	b.drawnTile = &tile
	b.turnPhase = placingTile
}

func (b *board) reject() {
	b.assertRoundPhase(tiling)
	b.assertTurnPhase(placingTile)
	b.openTiles[*b.drawnTile]++
	b.drawnTile = nil

	// TODO check if we should move to next round phase
	b.turnPhase = start
}

func (b *board) place(pos position, tile tile, farm bool) {
	b.assertRoundPhase(tiling)
	b.assertTurnPhase(placingTile)

	if b.turnPhase == start {
		b.openTiles[tile]--
	} else {
		b.drawnTile = nil
	}

	b.grid[pos] = tile

	// TODO check if we should move to next round phase

	if farm {
		b.farms[pos] = b.currentPlayer
	}

	b.nextPlayer()
}

func (b *board) claim(pos position) {
	b.assertRoundPhase(claiming)
	b.assertTurnPhase(start)
	b.claims[pos] = b.currentPlayer
	b.nextPlayer()
}

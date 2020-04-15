/*
Package contains the board state, how to calculate the next moves and execute them
A round begins with and initial board and ends with a score based on the number of claimed tiles per person
Each round has two phases, the tiling and claiming phase.

Each turn has three phases:

Start:
The player has to select an tile from the open tiles
or initiate the drawing of an closed tile.

Drawing tile:
The neutral player (the game it self) will select an random tile from the closed tiles

Placing tile:
The player has to place the tile in an valid position (posibly with an farm)

*/
package board

import "fmt"

type player int

const (
	one player = iota
	two
)

type turnPhase int

const (
	start turnPhase = iota
	drawingTile
	placingTile
)

type roundPhase int

const (
	tiling roundPhase = iota
	claiming
	end
)

const maxFarms = 4

// board contains the complete board state for this round.
type board struct {
	roundPhase    roundPhase
	turnPhase     turnPhase
	currentPlayer player

	// drawnTile contains the drawn (at random) tile to be placed by the next move.
	// This can be null when a tile has not been drawn
	drawnTile *tile

	// openTiles are the tiles placed face up
	// closedTiles are the tiles placed face down (in a pile)
	// Both contain the number of tiles per unique tile
	openTiles   map[tile]int
	closedTiles map[tile]int
	grid        map[position]tile
	farms       map[position]player
	claims      map[position]player

	// next* fields contains calculated data based on the other fields.
	// These make it more efficient to compute the next moves
	nextPlacements *placementTree
	nextClaims     map[player]map[position]bool
}

func new() (b *board) {
	b = &board{closedTiles: allTiles, grid: startGrid}
	for _, tile := range startGrid {
		b.closedTiles[tile]--
	}
	return b
}

func (b *board) numberOfFarms(p player) (count int) {
	for _, farmPlayer := range b.farms {
		if farmPlayer == p {
			count++
		}
	}
	return count
}

// Calculate all possible moves for the currrent board state
func (b *board) possibleMoves() (moves []moveExecuter) {
	switch b.roundPhase {
	case tiling:
		return b.possibleMovesWhileTiling()
	case claiming:
		return b.possibleMovesWhileClaiming()
	case end:
		return
	default:
		panic(fmt.Sprintf("Invalid round phase: %v", b.roundPhase))
	}
}

func (b *board) possibleMovesWhileTiling() (moves []moveExecuter) {
	switch b.turnPhase {
	case start:
		// The act of initiating the drawing and the drawing are separated.
		// The drawing is done by the neutral player (the game itself).
		// This is useful because it enables replaying and storing of game moves
		// without randomness involved.
		// It also allows for easy implementation of a game tree
		moves = append(moves, initiateDraw{})

		for tile, number := range b.openTiles {
			if number <= 0 {
				continue
			}

			for _, pos := range b.nextPlacements.positionsFor(tile) {
				if b.numberOfFarms(b.currentPlayer) < maxFarms {
					moves = append(moves, place{tile, pos, true})
				}
				moves = append(moves, place{tile, pos, false})
			}
		}
	case drawingTile:
		for tile, number := range b.closedTiles {
			for i := 0; i < number; i++ {
				moves = append(moves, draw{tile})
			}
		}
	case placingTile:
		for _, pos := range b.nextPlacements.positionsFor(*b.drawnTile) {
			if b.numberOfFarms(b.currentPlayer) < maxFarms {
				moves = append(moves, place{*b.drawnTile, pos, true})
			}
			moves = append(moves, place{*b.drawnTile, pos, false})
		}
		if len(moves) == 0 {
			moves = append(moves, reject{})
		}
	default:
		panic(fmt.Sprintf("Invalid turn phase: %v", b.turnPhase))
	}

	return moves
}

func (b *board) possibleMovesWhileClaiming() (moves []moveExecuter) {
	for pos, present := range b.nextClaims[b.currentPlayer] {
		if present {
			moves = append(moves, claim{pos})
		}
	}

	return moves
}

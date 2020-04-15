package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Check for each direction that you can move to a neighbore and come back to the same position
func TestNeighborsTurnAround(t *testing.T) {
	for dir := 0; dir < 6; dir++ {
		ldir := dir
		t.Run(string(dir), func(t *testing.T) {
			orig := position{0, 0, 0}
			neighPos := orig.neighbors()[ldir]
			newOrig := neighPos.neighbors()[(ldir+3)%6]
			assert.Equal(t, orig, newOrig)
		})
	}
}

// Check for each direction that you can move in a triangle and come back to the same position
func TestNeighborsTriangle(t *testing.T) {
	for dir := 0; dir < 6; dir++ {
		ldir := dir
		t.Run(string(dir), func(t *testing.T) {
			orig := position{0, 0, 0}
			neighPos := orig.neighbors()[ldir]
			neighPos = neighPos.neighbors()[(ldir+2)%6]
			neighPos = neighPos.neighbors()[(ldir+4)%6]
			assert.Equal(t, orig, neighPos)
		})
	}
}

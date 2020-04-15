package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that tiles do not contain uninitialized surfaces
// and the number of tiles is correct
func TestAllTilesNoAny(t *testing.T) {
	totalAmount := 0
	for tile, amount := range allTiles {
		for _, s := range tile {
			assert.NotEqual(t, any, s)
		}
		totalAmount += amount
	}
	assert.Equal(t, 40, totalAmount, "Number of tiles should be 40")
}

// Test that there are no rotations of tiles are in allTiles
func TestAllTilesNoRotations(t *testing.T) {
	for tile := range allTiles {
		rotTile := tile
		for i := 0; i < 5; i++ {
			rotTile = rotTile.rotRight()
			if tile == rotTile {
				// Skip if it is it's own rotation
				continue
			}
			assert.NotContains(t, allTiles, rotTile, "Rotation of a tile should not be present")
		}
	}
}

func TestStartGrid(t *testing.T) {
	for _, tile := range startGrid {
		assert.Contains(t, allTiles, tile, "Start tile should be present in start tiles")
	}
}

func TestRotation(t *testing.T) {
	tile := tile{land, water, hills, land, water, hills}
	rotTile := tile.rotRight().rotLeft()
	assert.Equal(t, tile, rotTile)

	rotTile = tile.rotLeft().rotRight()
	assert.Equal(t, tile, rotTile)

	rotTile = tile.rotRight().rotRight().rotRight()
	assert.Equal(t, tile, rotTile)
}

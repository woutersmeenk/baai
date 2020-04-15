package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositionForInitialGrid(t *testing.T) {
	tree := newPlacementTree(startGrid)
	actual := tree.positionsFor(tile{land, land, land, land, land, land})
	expected := []position{{-1, 0, 0}}
	assert.Equal(t, actual, expected)
}

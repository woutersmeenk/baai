package board

type placementTree struct {
	childs    map[tileSurface]*placementTree
	positions []position
}

// newPlacementTree creates a new tree
// initialGrid is the grid when the three startup tiles are placed
func newPlacementTree(initialGrid map[position]tile) (c *placementTree) {
	for pos := range initialGrid {
		for _, holePos := range pos.neighbors() {
			if _, ok := initialGrid[holePos]; !ok {
				continue
			}
			// We found a place for a tile (no tile present yet)
			var holeSurfaces [6]tileSurface
			for dir, neighOfHole := range holePos.neighbors() {
				// Check all tiles around the hole and see what the required surface is.
				if tile, ok := initialGrid[neighOfHole]; ok {
					// Fill in the two touching corners
					surfaceOne := tile[(dir+4)%6]
					surfaceTwo := tile[(dir+5)%6]
					holeSurfaces[dir] = surfaceOne
					holeSurfaces[(dir+1)%6] = surfaceTwo
				}
			}
		}
	}
	return c
}

// place adds the registers that the position is not available anymore.
// Also add the new places where tiles can be placed.
// grid is the curent grid before the placement, this will not be modified
func (c *placementTree) place(tile tile, pos position, grid map[position]tile) {
}

func (c *placementTree) positionsFor(tile tile) (result []position) {
	currentTile := tile
	for r := 0; r < 6; r++ {
		// TODO rotate results
		result = append(result, c.possibleForCorners(currentTile[:])...)
		currentTile = currentTile.rotRight()
	}
	return result
}

func (c *placementTree) possibleForCorners(corners []tileSurface) (result []position) {
	if c == nil {
		return
	}

	// TODO not efficient?
	result = append(result, c.positions...)
	nextNode := c.childs[corners[0]]
	result = append(result, nextNode.possibleForCorners(corners[1:])...)

	return result
}

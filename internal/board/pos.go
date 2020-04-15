package board

// position represents the position in the hex grid of tiles
// We use axial coordinates as explained here:
// https://www.redblobgames.com/grids/hexagons/#coordinates
// rot is the number of times the tile rotated right
type position struct {
	q, r int
	rot  int
}

func (pos position) neighbors() [6]position {
	return [6]position{
		{pos.q - 1, pos.r + 0, pos.rot},
		{pos.q + 0, pos.r - 1, pos.rot},
		{pos.q + 1, pos.r - 1, pos.rot},
		{pos.q + 1, pos.r + 0, pos.rot},
		{pos.q + 0, pos.r + 1, pos.rot},
		{pos.q - 1, pos.r + 1, pos.rot},
	}
}

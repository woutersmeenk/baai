package board

import "fmt"

type tileSurface int

const (
	any tileSurface = iota
	land
	water
	hills
)

func (s tileSurface) String() string {
	switch s {
	case any:
		return "any"
	case land:
		return "land"
	case water:
		return "water"
	case hills:
		return "hills"
	default:
		panic(fmt.Sprintf("Invalid tileSurface %v", int(s)))
	}
}

// tile describes what a hex tile looks like.
// it describes the surface of the six corners of the tile.
//    1   2
//    /---\
// 0 /     \ 3
//   \     /
//    \___/
//    5    4
// Starting with the left most point moving clockwise around.
// Only the tiles in allTiles are conical. All other permutations of the tiles are not valid.
type tile [6]tileSurface

func (t tile) rotRight() (result tile) {
	copy(result[1:6], t[0:5])
	result[0] = t[5]
	return
}
func (t tile) rotLeft() (result tile) {
	copy(result[0:5], t[1:6])
	result[5] = t[0]
	return
}

var allTiles = map[tile]int{
	[6]tileSurface{water, water, water, water, land, land}:   4,
	[6]tileSurface{land, water, water, water, land, land}:    5,
	[6]tileSurface{land, land, land, land, land, land}:       3,
	[6]tileSurface{land, land, land, land, land, water}:      1,
	[6]tileSurface{hills, hills, land, hills, hills, land}:   1,
	[6]tileSurface{land, land, land, hills, hills, hills}:    3,
	[6]tileSurface{hills, hills, land, land, land, land}:     5,
	[6]tileSurface{hills, hills, land, land, hills, land}:    1,
	[6]tileSurface{hills, hills, hills, hills, hills, hills}: 1,
	[6]tileSurface{land, land, land, land, land, hills}:      3,
	[6]tileSurface{land, land, hills, hills, hills, hills}:   2,
	[6]tileSurface{hills, hills, hills, hills, land, hills}:  1,
	[6]tileSurface{hills, hills, land, land, water, land}:    2,
	[6]tileSurface{water, water, water, water, land, water}:  1,
	[6]tileSurface{land, water, water, water, land, hills}:   2,
	[6]tileSurface{land, land, land, land, water, water}:     1,
	[6]tileSurface{land, hills, land, land, water, land}:     1,
	[6]tileSurface{hills, land, land, land, hills, land}:     1,
	[6]tileSurface{land, land, hills, land, water, water}:    1,
	[6]tileSurface{land, hills, hills, land, water, water}:   1,
}

var startGrid = map[position]tile{
	{0, 0, 0}:  [6]tileSurface{land, land, land, land, land, land},
	{0, -1, 2}: [6]tileSurface{land, land, land, land, land, land},
	{1, -1, 2}: [6]tileSurface{hills, hills, land, land, water, land},
}

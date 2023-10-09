package roomgrowinggenerator

// VAULT DEFINITION SYNTAX:
const (
	charAny   = ' ' // can be any tile
	charWall  = '#'
	charFence = '\'' // Can be seen through, but not walkable
	charFloor = '.'  // Causes to create new room id if neccessary (i.e. if a vault contains a door)
	//                  ...so it does NOT guarantee to change room ID.
	charFloorOldId = ',' // Like Floor, but forces to retain old room id (useful for internal vaults' outer borders).
	//                         ...This tile will be treated by doors gen as an "outer room" floor, not as a part of the vault.
	//                         ...Ignoring this tile may cause redundant/illogical doors to be placed.
	//                         ...It's also redundant for vaults without doors, but will work anyway
	charDoor = '+'
)

// Vaults that may be placed at the start of the generation as the first rooms.
// Restrictions: no floor borders and no doors on the borders.
var initialVaults = [][]string{
	{
		"#############",
		"#.....+.....#",
		"#.....'.....#",
		"#.....'.....#",
		"#+'''''''''+#",
		"#.....'.....#",
		"#.....'.....#",
		"#.....+.....#",
		"#############",
	},
	{
		"   #######   ",
		"   #.....#   ",
		"   #.....#   ",
		"####.....####",
		"#...........#",
		"#...........#",
		"#...........#",
		"#...........#",
		"####.....####",
		"   #.....#   ",
		"   #.....#   ",
		"   #######   ",
	},
	{
		"   #######   ",
		"   #.....#   ",
		"   #..#..#   ",
		"####..#..####",
		"#.....#.....#",
		"#.#########.#",
		"#.....#.....#",
		"####..#..####",
		"   #..#..#   ",
		"   #.....#   ",
		"   #######   ",
	},
}

// Those vaults are treated as rooms themselves (so are placed OUTSIDE other rooms, appended to them).
// Outside vaults SHOULD have at least one outer door.
// TODO: randomly add a door if it does not have any?
var outsideVaults = [][]string{
	{
		"##########",
		"#........+",
		"#.########",
		"#........#",
		"########.#",
		"#........#",
		"##########",
	},
	{
		"#####    ",
		"#...#    ",
		"#...#    ",
		"#...#####",
		"#.......#",
		"#.......+",
		"#########",
	},
	{
		"#####    ",
		"#...#    ",
		"#...#    ",
		"#...#####",
		"#.......#",
		"#.......#",
		"#.......#",
		"#####...#",
		"    #...#",
		"    #...#",
		"    ##+##",
	},
	{
		"  #####  ",
		" ##...## ",
		"##.....##",
		"#.......+",
		"##.....##",
		" ##...## ",
		"  #####  ",
	},
	{
		"   ####   ",
		"   #..#   ",
		"   #..#   ",
		"####..####",
		"#........+",
		"#........+",
		"####..####",
		"   #..#   ",
		"   #..#   ",
		"   ####   ",
	},
	{
		"######     ",
		"#....#     ",
		"#....#     ",
		"#....######",
		"#....'....#",
		"#....+....+",
		"#....'....#",
		"#....######",
		"#....#     ",
		"#....#     ",
		"######     ",
	},
}

// Vaults placed INSIDE other rooms
// Inside vaults should be surrounded by floor tiles to enforce passability.
// Ignoring this will cause unconnected maps generation
var insideVaults = [][]string{
	{
		"...",
		".#.",
		"...",
	},
	{
		"....",
		".##.",
		".##.",
		"....",
	},
	{
		".........",
		".#.#.#.#.",
		".........",
	},
	{
		" ... ",
		"..#..",
		".###.",
		"..#..",
		" ... ",
	},
	{
		",,,,,,,,",
		",######,",
		",#....#,",
		",#....#,",
		",#....+,",
		",#....#,",
		",######,",
		",,,,,,,,",
	},
	{
		" ,,,,,, ",
		",,####,,",
		",##..##,",
		",#....#,",
		",#....+,",
		",##..##,",
		",,####,,",
		" ,,,,,, ",
	},
	{
		",,,,,,,,",
		",'''''',",
		",'....',",
		",'....',",
		",'....+,",
		",'....',",
		",'''''',",
		",,,,,,,,",
	},
	{
		"........",
		".######.",
		".#..#...",
		".#..#...",
		".#......",
		".#..#...",
		".#..#...",
		".######.",
		"........",
	},
	{
		".........",
		".#######.",
		"....#....",
		"....#....",
		".........",
		"....#....",
		"....#....",
		".#######.",
		".........",
	},
}
